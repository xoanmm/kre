import React from 'react';
import styles from './VerifyEmail.module.scss';


function VerifyEmail() {
  return (
    <div className={ styles.bg }>
      <div className={ styles.grid }>
        <div className={ styles.container }>
          <h1>Login link sent</h1>
          <p className={ styles.subtitle }>
            An email with an access link has been sent to your email address.
          </p>
        </div>
      </div>
    </div>
  );
}

export default VerifyEmail;
