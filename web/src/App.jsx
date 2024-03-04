import React from "react";
import RouteRules from "./routes/RouteRules.jsx";
import {BrowserRouter} from "react-router-dom";

const App = () => {
  return (
    <BrowserRouter>
      <RouteRules />
    </BrowserRouter>
  );
};

export default App;
