import React from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

export default function LandingPage() {
  return (
    <>
      <section className="text-center mb-12">
        <h2 className="text-4xl font-bold text-gray-900 mb-4">
          Welcome to API Diff Tool
        </h2>
        <p className="text-xl text-gray-600 max-w-2xl mx-auto">
          Easily compare and analyze differences between API responses with our
          powerful diff viewer.
        </p>
      </section>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-12">
        <Card>
          <CardHeader>
            <CardTitle>Easy Comparison</CardTitle>
          </CardHeader>
          <CardContent>
            <p>
              Compare API responses side-by-side with our intuitive diff viewer.
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>JSON Support</CardTitle>
          </CardHeader>
          <CardContent>
            <p>
              Full support for JSON data structures, making it easy to spot
              changes in complex API responses.
            </p>
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Syntax Highlighting</CardTitle>
          </CardHeader>
          <CardContent>
            <p>
              Benefit from syntax highlighting to easily identify changes in
              your API responses.
            </p>
          </CardContent>
        </Card>
      </div>

      <div className="text-center">
        <Button size="lg" asChild>
          <Link to="/diff">Try API Diff Viewer Now</Link>
        </Button>
      </div>
    </>
  );
}
