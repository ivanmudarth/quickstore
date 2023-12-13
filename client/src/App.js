import { ChakraProvider } from "@chakra-ui/react";
import Home from "./components/Home";
import TestTabs from "./components/shadcn";
import { Calendar } from "./components/ui/calendar";
import React, { useState } from "react";
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
  const [date, setDate] = useState(new Date());
  return (
    <div>
      <ChakraProvider>
        <Home />
      </ChakraProvider>
      <TestTabs />
      <Calendar
        mode="single"
        selected={date}
        onSelect={setDate}
        className="rounded-md border"
      />
    </div>
  );
}

export default App;
