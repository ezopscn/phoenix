import React from "react";
import ForbiddenErrorImage from "../../assets/image/error/403.svg";

const ForbiddenError = () => {
  return (
    <>
      <div className="error-code">403</div>
      <img src={ForbiddenErrorImage} alt="" />
    </>
  );
};

export default ForbiddenError;
