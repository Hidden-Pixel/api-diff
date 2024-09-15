import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Layout from "./components/Layout";
import LandingPage from "./pages/LandingPage";
import ApiDiffViewer from "./pages/ApiDiffViewer";
import ApiForm from "./pages/ApiForm";

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/diff" element={<ApiDiffViewer />} />
          <Route path="/form" element={<ApiForm />} />
        </Routes>
      </Layout>
    </Router>
  );
}

export default App;
