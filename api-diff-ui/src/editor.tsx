import { MonacoDiffEditor } from "react-monaco-editor";

export default function Editor() {
  const code1 = "// your original code...";
  const code2 = "// a different version...";
  const options = {
    renderSideBySide: true,
  };
  return (
    <MonacoDiffEditor
      width="1200"
      height="800"
      theme="vs-dark"
      language="javascript"
      original={code1}
      value={code2}
      options={options}
    />
  );
}
