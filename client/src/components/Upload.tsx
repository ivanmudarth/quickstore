import React, { useState } from "react";

import { Card, CardDescription, CardHeader, CardTitle } from "./ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";
import { Flex } from "@chakra-ui/react";
import FileUpload from "./FileUpload";
import UrlUpload from "./UrlUpload";
import { FinishUploadProp } from "./utils/sharedTypes";

enum uploadTypes {
  "File",
  "Url",
}

function Upload(props: FinishUploadProp) {
  const [uploadTypeSwitch, setUploadTypeSwitch] = useState<{
    uploadTypeSwitch: uploadTypes;
  }>({ uploadTypeSwitch: uploadTypes.File });

  function handleUploadTypeSwitch(value: string) {
    var newType = uploadTypes.Url;
    if (uploadTypeSwitch.uploadTypeSwitch == uploadTypes.Url) {
      newType = uploadTypes.File;
    }
    setUploadTypeSwitch({ uploadTypeSwitch: newType });
  }

  return (
    <Card>
      <CardHeader>
        <Flex justifyContent="space-between" alignItems="center" py={2}>
          <CardTitle>Upload</CardTitle>
          <Select onValueChange={handleUploadTypeSwitch}>
            <SelectTrigger className="w-[130px]">
              <SelectValue placeholder="Upload Type" />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="file">File</SelectItem>
              <SelectItem value="url">URL</SelectItem>
            </SelectContent>
          </Select>
        </Flex>

        <CardDescription>
          Upload a file or URL here. Make them searchable by adding your own
          tags.
        </CardDescription>
      </CardHeader>
      {uploadTypeSwitch.uploadTypeSwitch === uploadTypes.File ? (
        <FileUpload finishUpload={props.finishUpload} />
      ) : (
        <UrlUpload finishUpload={props.finishUpload} />
      )}
    </Card>
  );
}

export default Upload;
