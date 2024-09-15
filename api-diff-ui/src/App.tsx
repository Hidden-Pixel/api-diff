import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import LandingPage from "./pages/LandingPage";
import ApiDiffViewer from "./pages/ApiDiffViewer";

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/diff" element={<ApiDiffViewer />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;
