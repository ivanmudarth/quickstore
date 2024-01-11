import { UrlInfo } from "./utils/sharedTypes";

interface Props {
  urlinfo: UrlInfo;
}

function Url(props: Props) {
  return (
    <div
      className={"space-y-3 w-[275px] rounded-md border"}
      style={{ padding: "15px" }}
      {...props}
    >
      <div className="overflow-hidden">
        <a href={props.urlinfo.URL} target="_blank">
          <img
            src={props.urlinfo.ImageURL}
            alt={props.urlinfo.Title}
            className="rounded-md"
            style={{ maxHeight: "150px" }}
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
