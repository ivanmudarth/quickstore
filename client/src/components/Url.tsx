import { UrlInfo } from "./utils/sharedTypes";

interface Props {
  urlinfo: UrlInfo;
}

function Url(props: Props) {
  return (
    <div
      className={"space-y-3 w-[275px] rounded-md border"}
      style={{ padding: "15px", boxShadow: "0 2px 4px rgba(0, 0, 0, 0.1)" }}
      {...props}
    >
      <div className="overflow-hidden">
        <a href={props.urlinfo.URL} target="_blank">
          <img
            src={props.urlinfo.ImageURL}
            alt={props.urlinfo.Title}
            className="rounded-md"
            style={{
              width: "1024px",
              height: "2500px",
              objectFit: "cover", // This property will crop the image to fit the specified dimensions
              objectPosition: "top left",
              maxHeight: "150px", // This is to set a maximum height if needed
            }}
          />
        </a>
      </div>
      <div className="space-y-1 text-sm">
        <h3 className="font-medium leading-none" style={{ fontSize: "medium" }}>
          {props.urlinfo.Title}
        </h3>
        <p
          className="text-s text-muted-foreground" // TODO: add < if size is less than 0.01 MB
        >
          {props.urlinfo.URL}
        </p>
        <p className="text-s text-muted-foreground">
          <u>User Tags:</u> {props.urlinfo.UserTags?.join(", ")}
        </p>
        <p className="text-s text-muted-foreground">
          <u>Auto Tags:</u> {props.urlinfo.AutoTags?.join(", ")}
        </p>
      </div>
    </div>
  );
}

export default Url;
