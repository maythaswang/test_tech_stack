import { useState } from "react";
import reactLogo from "./assets/images/react.svg";
import viteLogo from "/vite.svg";
import "./assets/css/tailwind.css";
import "./assets/css/App.css";
import "./assets/css/index.css";


function Index() {
  const [count, setCount] = useState(0);

  return (
    <>
      <main>
        <div>
          <a href="https://vite.dev" target="_blank">
            <img src={viteLogo} className="logo" alt="Vite logo" />
          </a>
          <a href="https://react.dev" target="_blank">
            <img src={reactLogo} className="logo react" alt="React logo" />
          </a>
        </div>
        <h1>Vite + React</h1>

        <div className="card">
          <button onClick={() => setCount((count) => count + 1)}>
            count is {count}
          </button>
          <br />
          <p>
            Edit <code>src/App.tsx</code> and save to test HMR
          </p>
        </div>
        <p className="read-the-docs">
          Click on the Vite and React logos to learn more
        </p>
      </main>
    </>
  );
}

export default Index;
