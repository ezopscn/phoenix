import React from 'react';
import RouteRules from './routes/RouteRules.jsx';
import { HashRouter } from 'react-router-dom';
import { AuthRouter } from './routes/RouteMatch.jsx';

const App = () => {
  return (
    <HashRouter>
      <AuthRouter>
        <RouteRules />
      </AuthRouter>
    </HashRouter>
  );
};

export default App;
