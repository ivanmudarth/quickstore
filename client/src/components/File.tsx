import { FileInfo } from "./utils/sharedTypes";
// @ts-ignore
import pdfIcon from "../public/pdf_icon.png";
import { Tag } from "@chakra-ui/react";

interface Props {
  fileinfo: FileInfo;
}

function File(props: Props) {
  return (
    <div
      className={"space-y-3 w-[275px] rounded-md border"}
      style={{ padding: "15px", boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)" }}
      {...props}
    >
      <div className="overflow-hidden">
        {props.fileinfo.Type == "Image" ? (
          <img
            src={props.fileinfo.FileURL}
            alt={props.fileinfo.Name}
            className="rounded-md"
            // style={{ maxWidth: "300px", maxHeight: "300px" }}
          />
        ) : (
          <a
            href={props.fileinfo.FileURL}
            // download={encodeURIComponent(props.info.Name)} TODO: doesn't work
          >
            <img
              src={pdfIcon}
              alt="PDF Logo"
              className="rounded-md"
              style={{ maxHeight: "150px" }}
            />
          </a>
        )}
      </div>
      <div className="space-y-1 text-sm">
        <h3 className="font-medium leading-none" style={{ fontSize: "medium" }}>
          {props.fileinfo.Name}
        </h3>
        <p
          className="text-s text-muted-foreground" // TODO: add < if size is less than 0.01 MB
          style={{ marginBottom: "8px" }}
        >
          {props.fileinfo.Size} MB
        </p>
        {props.fileinfo.UserTags?.map((tag, _) => (
          <Tag size={"sm"} style={{ marginRight: "10px" }}>
            {tag}
          </Tag>
        ))}
        {props.fileinfo.AutoTags?.map((tag, _) => (
          <Tag size={"sm"} color={"green"} style={{ marginRight: "10px" }}>
            {tag}
          </Tag>
        ))}
      </div>
    </div>
  );
}

export default File;
