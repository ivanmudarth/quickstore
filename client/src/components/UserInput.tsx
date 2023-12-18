import React from "react";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "./ui/tabs";

interface Props {
  components: {
    uploadComponent: JSX.Element;
    searchComponent: JSX.Element;
  };
}

const UserInput: React.FC<Props> = ({ components }) => {
  const { uploadComponent: UploadComponent, searchComponent: SearchComponent } =
    components;

  return (
    <div style={{ paddingTop: "50px" }}>
      <Tabs defaultValue="upload" className="w-[500px]">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="upload">Upload</TabsTrigger>
          <TabsTrigger value="search">Search</TabsTrigger>
        </TabsList>
        <TabsContent value="upload">{UploadComponent}</TabsContent>
        <TabsContent value="search">{SearchComponent}</TabsContent>
      </Tabs>
    </div>
  );
};

export default UserInput;
