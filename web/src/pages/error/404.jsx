import React from 'react';
import NotFoundErrorImage from '../../assets/image/error/404.svg';

const NotFoundError = () => {
  return (
    <>
      <div className="error-code">404</div>
      <img src={NotFoundErrorImage} alt="" />
    </>
  );
};

export default NotFoundError;
