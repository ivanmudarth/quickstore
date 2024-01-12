import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Button } from "./ui/button";
import { CardContent, CardFooter } from "./ui/card";
import axios from "axios";
import { processTagInput } from "./utils/TagInput";
import React, { useState } from "react";
import { FinishUploadProp } from "./utils/sharedTypes";
import { Spinner } from "@chakra-ui/react";

function FileUpload(props: FinishUploadProp) {
  const [file, setFile] = useState<{ file: File | null }>({ file: null });
  const [rawTagInput, setTagInput] = useState<{ rawTagInput: string }>({
    rawTagInput: "",
  });
  const [isLoading, setIsLoading] = useState<{ isLoading: Boolean }>({
    isLoading: false,
  });

  // TODO: input requirements
  // username must be unique (check)
  // filename, each tag must be less than / equal to 15 chars

  function handleFileChange(event: React.ChangeEvent<HTMLInputElement>) {
    const selectedFile = (event.target as HTMLInputElement).files?.[0] || null;
    setFile({ file: selectedFile });
  }

  function handleTagChange(event: React.ChangeEvent<HTMLInputElement>) {
    setTagInput({ rawTagInput: event.target.value });
  }

  function handleUpload(event: React.FormEvent<HTMLFormElement>) {
    // TODO:
    // limit number of files to add
    // case where no file is selected
    // limit type of file to upload
    // upload progress bar / disable input/submit while uploading
    // display error message if error in upload
    // user generated tags
    // clear input values on submit
    // ensure user tag input is valid (no invalid chars)

    // restrict file type to pdf and images only

    event.preventDefault();

    const formData = new FormData();
    const userTags = processTagInput(rawTagInput.rawTagInput);

    formData.append("file", file.file ?? "");
    formData.append("fileName", file.file?.name ?? "");
    userTags.forEach((tag) => {
      formData.append("tags[]", tag);
    });

    const url = "http://localhost:8080/upload_file";
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
          <Label htmlFor="file">File:</Label>
          <Input type="file" onChange={handleFileChange} />
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

export default FileUpload;
