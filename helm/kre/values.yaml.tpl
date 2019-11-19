# Default values for konstellation Runtime Environment.
mongodb:
  service:
    name: "mongodb"
  mongodbDatabase: "localKRE"
  mongodbUsername: "admin"

config:
  admin:
    apiAddress: ":80"
    frontendBaseURL: "http://localhost:3000"
    corsEnabled: true
  smtp:
    enabled: false
  auth:
    verificationCodeDurationInMinutes: 1
    sessionDurationInHours: 1
    jwtSignSecret: jwt_secret
    secureCookie: false
  mongodb:
    address: "mongodb://mongodb:27017"
    dbName: "localKRE"

adminApi:
  image:
    repository: konstellation/kre-admin-api
    tag: ${ADMIN_API_IMAGE_TAG}
    pullPolicy: Always
  service:
    port: 4000
  tls:
    enabled: false
    host: admin-api-kre.local
