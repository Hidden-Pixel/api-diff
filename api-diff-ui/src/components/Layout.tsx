import React from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";

export default function Layout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen flex flex-col bg-gradient-to-b from-gray-100 to-gray-200">
      <header className="bg-white shadow-sm">
        <nav className="container mx-auto px-4 py-4 flex justify-between items-center">
          <Link to="/" className="text-2xl font-bold text-gray-800">
            API Diff Tool
          </Link>
          <div className="space-x-4">
            <Button variant="ghost" asChild>
              <Link to="/">Home</Link>
            </Button>
            <Button variant="ghost" asChild>
              <Link to="/diff">Diff Viewer</Link>
            </Button>
            <Button variant="ghost" asChild>
              <Link to="/form">Data Entry</Link>
            </Button>
          </div>
        </nav>
      </header>

      <main className="flex-grow container mx-auto px-4 py-8">{children}</main>

      <footer className="bg-gray-800 text-white py-8">
        <div className="container mx-auto px-4 text-center">
          <p>&copy; 2023 API Diff Tool. All rights reserved.</p>
        </div>
      </footer>
    </div>
  );
}
