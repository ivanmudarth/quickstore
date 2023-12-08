import React, { useState } from "react";
import { Switch, HStack, Text } from "@chakra-ui/react";

// TODO: add TagSearch component
function UserInput({ uploadComponent, searchComponent }) {
  const [isUpload, setIsUpload] = useState(true);

  return (
    <>
      <HStack>
        <Text>Upload or Search</Text>
        <Switch onChange={() => setIsUpload(!isUpload)}></Switch>
      </HStack>
      {isUpload ? uploadComponent : searchComponent}
    </>
  );
}

export default UserInput;
