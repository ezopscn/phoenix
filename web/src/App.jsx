import React from "react";
import { BrowserRouter } from "react-router-dom";
import RouteRules from "./routes/RouteRules.jsx";

const App = () => {
  return (
    <BrowserRouter>
      <RouteRules />
    </BrowserRouter>
  );
};

export default App;
