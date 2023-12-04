import "./App.css";
import { ChakraProvider } from "@chakra-ui/react";
import Home from "./components/Home";

/*
App
  Home
    UserInput
      UploadFile - request /upload
      TagSearch - request /search
    FileDisplay - request /display
      File(fileInfo)
*/

function App() {
  return (
    <ChakraProvider>
      <Home />
    </ChakraProvider>
  );
}

export default App;
