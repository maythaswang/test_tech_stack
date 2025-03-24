import Header from "./assets/components/Header";
import "./assets/css/tailwind.css";
import "./assets/css/App.css";
import "./assets/css/index.css";

import AppRouter from "./AppRouter";
import { BrowserRouter as Router } from "react-router-dom";

function App() {
  return (
    <>
      <header>
        <Router>
          <Header />
          <AppRouter />
        </Router>
      </header>
    </>
  );
}

export default App;
