import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Button } from "./ui/button";
import { CardContent, CardFooter } from "./ui/card";
import axios from "axios";
import { processTagInput } from "./utils/TagInput";
import React, { useState } from "react";
import { FinishUploadProp } from "./utils/sharedTypes";
import { Spinner } from "@chakra-ui/react";

function UrlUpload(props: FinishUploadProp) {
  const [urlInput, setUrlInput] = useState<{ urlInput: string }>({
    urlInput: "",
  });
  const [rawTagInput, setTagInput] = useState<{ rawTagInput: string }>({
    rawTagInput: "",
  });
  const [isLoading, setIsLoading] = useState<{ isLoading: Boolean }>({
    isLoading: false,
  });

  function handleTagChange(event: React.ChangeEvent<HTMLInputElement>) {
    setTagInput({ rawTagInput: event.target.value });
  }

  function handleUrlChange(event: React.ChangeEvent<HTMLInputElement>) {
    setUrlInput({ urlInput: event.target.value });
  }

  function handleUpload(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    const formData = new FormData();
    const userTags = processTagInput(rawTagInput.rawTagInput);

    formData.append("url", urlInput.urlInput);
    userTags.forEach((tag) => {
      formData.append("tags[]", tag);
    });
    const url = "http://localhost:8080/upload_url";
    const config = { headers: { "content-type": "multipart/form-data" } };
    setIsLoading({ isLoading: true });
    axios
      .post(url, formData, config)
      .then((response) => {
        console.log(response.data);
        setIsLoading({ isLoading: false });
        props.finishUpload();
      })
      .catch((error) => {
        console.log(error);
      });
  }

  return (
    <form onSubmit={handleUpload}>
      <CardContent className="space-y-2">
        <div className="space-y-1">
          <Label htmlFor="url">URL:</Label>
          <Input
            placeholder="Enter a valid website URL"
            onChange={handleUrlChange}
          ></Input>
        </div>
        <div className="space-y-1">
          <Label htmlFor="tags">Tags:</Label>
          <Input
            placeholder="Enter a comma separated list"
            onChange={handleTagChange}
          ></Input>
        </div>
      </CardContent>
      <CardFooter>
        <Button
          type="submit"
          disabled={!!isLoading.isLoading}
          style={{ marginRight: "10px" }}
        >
          Upload
        </Button>
        {isLoading.isLoading && <Spinner />}
      </CardFooter>
    </form>
  );
}

export default UrlUpload;
