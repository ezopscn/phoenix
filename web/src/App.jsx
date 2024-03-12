import React from "react";
import RouteRules from "./routes/RouteRules.jsx";
import { BrowserRouter } from "react-router-dom";
import { AuthRouter } from "./routes/RouteMatch.jsx";

const App = () => {
  return (
    <BrowserRouter>
      <AuthRouter>
        <RouteRules />
      </AuthRouter>
    </BrowserRouter>
  );
};

export default App;
