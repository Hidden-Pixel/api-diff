import React, { useState } from "react";
import { Editor, DiffEditor } from "@monaco-editor/react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "@/components/ui/button";

export default function ApiDiffViewer() {
  const [oldData, setOldData] = useState("");
  const [newData, setNewData] = useState("");
  const [showDiff, setShowDiff] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    setShowDiff(true);
  };

  const editorOptions = {
    minimap: { enabled: false },
    scrollBeyondLastLine: false,
    fontSize: 14,
    lineNumbers: "on",
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-6">API Diff Viewer</h1>
      <form onSubmit={handleSubmit} className="mb-6">
        <Card>
          <CardHeader>
            <CardTitle>API Data</CardTitle>
          </CardHeader>
          <CardContent>
            <Tabs defaultValue="old" className="w-full">
              <TabsList className="grid w-full grid-cols-2">
                <TabsTrigger value="old">Old Data</TabsTrigger>
                <TabsTrigger value="new">New Data</TabsTrigger>
              </TabsList>
              <TabsContent value="old">
                <Editor
                  height="300px"
                  language="json"
                  value={oldData}
                  onChange={(value) => setOldData(value || "")}
                  options={editorOptions}
                  className="border rounded-md"
                />
              </TabsContent>
              <TabsContent value="new">
                <Editor
                  height="300px"
                  language="json"
                  value={newData}
                  onChange={(value) => setNewData(value || "")}
                  options={editorOptions}
                  className="border rounded-md"
                />
              </TabsContent>
            </Tabs>
          </CardContent>
        </Card>
        <Button type="submit" className="mt-4">
          Compare
        </Button>
      </form>
      {showDiff && (
        <Card>
          <CardHeader>
            <CardTitle>API Diff</CardTitle>
          </CardHeader>
          <CardContent>
            <DiffEditor
              height="400px"
              language="json"
              original={oldData}
              modified={newData}
              options={{
                ...editorOptions,
                readOnly: true,
                renderSideBySide: true,
              }}
              className="border rounded-md"
            />
          </CardContent>
        </Card>
      )}
    </div>
  );
}
