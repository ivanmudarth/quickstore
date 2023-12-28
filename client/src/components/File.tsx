// @ts-ignore
import pdfIcon from "../public/PDF_icon.png";

interface Props {
  info: {
    URL: string;
    Name: string;
    Size: string;
    Type: string;
    UserTags: string[];
    AutoTags: string[];
  };
}

function File(props: Props) {
  return (
    <div
      className={"space-y-3 w-[275px] rounded-md border"}
      style={{ padding: "15px" }}
      {...props}
    >
      <div className="overflow-hidden">
        {props.info.Type == "Image" ? (
          <img
            src={props.info.URL}
            alt={props.info.Name}
            className="rounded-md"
            // style={{ maxWidth: "300px", maxHeight: "300px" }}
          />
        ) : (
          <a
            href={props.info.URL}
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
          {props.info.Name}
        </h3>
        <p className="text-s text-muted-foreground">{props.info.Size} MB</p>
        <p className="text-s text-muted-foreground">
          <u>User Tags:</u> {props.info.UserTags?.join(", ")}
        </p>
        <p className="text-s text-muted-foreground">
          <u>Auto Tags:</u> {props.info.AutoTags?.join(", ")}
        </p>
      </div>
    </div>
  );
}

export default File;
