import logging
import os
import asyncio
import json

from nats.aio.client import Client as NATS

from grpclib.utils import graceful_exit
from grpclib.server import Server, Stream
from grpclib.reflection.service import ServerReflection
from protobuf_to_dict import protobuf_to_dict

# generated by protoc
from public_input_pb2 import Request, Response
from public_input_grpc import EntrypointBase

nc = NATS()

with open(os.environ['KRT_NATS_SUBJECTS_FILE']) as json_file:
    subjects = json.load(json_file)
    print(f"Loaded NATS subject file: {subjects}")


def MessageToDict(message):
    messageDict = {}

    for descriptor in message.DESCRIPTOR.fields:
        key = descriptor.name
        value = getattr(message, descriptor.name)

        if descriptor.label == descriptor.LABEL_REPEATED:
            messageList = []

            for subMessage in value:
                if descriptor.type == descriptor.TYPE_MESSAGE:
                    messageList.append(MessageToDict(subMessage))
                else:
                    messageList.append(subMessage)

            messageDict[key] = messageList
        else:
            if descriptor.type == descriptor.TYPE_MESSAGE:
                messageDict[key] = MessageToDict(value)
            else:
                messageDict[key] = value

    return messageDict


class Result:
    def __init__(self, reply='', data=None, error=None):
        self.reply = reply
        self.data = data
        self.error = error

    def from_nats_msg(self, msg):
        try:
            data = json.loads(msg.data.decode())
        except Exception as err:
            raise Exception(
                f"error parsing msg.data because is not a valid JSON: {str(err)}")
        self.reply = data.get("reply")
        self.data = data.get("data")
        self.error = data.get("error")

    def to_json(self):
        return bytes(json.dumps(self.__dict__), encoding='utf-8')


class Entrypoint(EntrypointBase):

    async def Greet(self, stream: Stream[Request, Response]) -> None:
        request = await stream.recv_message()
        assert request is not None
        try:
            request_dict = MessageToDict(request)
            print(f'gRPC message received: {json.dumps(request_dict)}')

            result = Result(data=request_dict)

            nats_subject = subjects['Greet']
            print(f"Request message to NATS subject {nats_subject}")
            res = await nc.request(nats_subject, result.to_json(), timeout=60)
            response = Result()
            response.from_nats_msg(res)
            print(f"NATS message received: {response.to_json()}")

            await stream.send_message(Response(greeting=response.data['greeting']))
            print(
                f'gRPC successfully response: {str(response.data)}')
        except Exception as err:
            logging.error(f'Exception on gRPC call : {err})')
            await stream.send_message(Response(greeting=f'Exception on gRPC call: {err})'))


async def run(loop, host: str = '0.0.0.0', port: int = 9000) -> None:
    print(f"Connecting to NATS {os.environ['KRT_NATS_SERVER']}...")
    await nc.connect(os.environ['KRT_NATS_SERVER'], loop=loop)

    services = ServerReflection.extend([Entrypoint()])

    server = Server(services)
    with graceful_exit([server]):
        await server.start(host, port)
        print(f'Serving gPRC server on {host}:{port}')
        await server.wait_closed()


if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    asyncio.ensure_future(run(loop))
    loop.run_forever()
