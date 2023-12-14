import React, { useState, ComponentType } from "react";
import { Switch, HStack, Text } from "@chakra-ui/react";

interface Props {
  components: {
    uploadComponent: JSX.Element;
    searchComponent: JSX.Element;
  };
}

// TODO: add TagSearch component
const UserInput: React.FC<Props> = ({ components }) => {
  const [isUpload, setIsUpload] = useState(true);
  const { uploadComponent: UploadComponent, searchComponent: SearchComponent } =
    components;

  return (
    <>
      <HStack>
        <Text>Upload or Search</Text>
        <Switch onChange={() => setIsUpload(!isUpload)}></Switch>
      </HStack>
      {isUpload ? UploadComponent : SearchComponent}
    </>
  );
};

export default UserInput;
