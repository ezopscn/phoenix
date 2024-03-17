import React from 'react';
import '../../assets/css/error.less';
import { FooterText, LogoWithTitleNoBg } from '../../common/Resource.jsx';
import { Outlet, useNavigate } from 'react-router';
import { Button } from 'antd';

const ErrorLayout = () => {
  const navigate = useNavigate();
  return (
    <>
      <div className="error-container">
        <div className="error-header">
          <img src={LogoWithTitleNoBg} alt="" draggable="false" />
        </div>
        <div className="error-body">
          <div className="error-info">
            <div>
              <Outlet />
            </div>
            <div>
              <Button onClick={() => navigate('/')} className="error-btn" type="primary">
                返回首页
              </Button>
            </div>
          </div>
        </div>
        <div className="error-footer">
          <FooterText />
        </div>
      </div>
    </>
  );
};

export default ErrorLayout;
