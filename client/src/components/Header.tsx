// @ts-ignore
import logo from "../public/logo.png";

function Header() {
  return (
    // <header className="sticky top-0 z-50 w-full border-b bg-background/95 backdrop-blur supports-[backdrop-filter]:bg-background/60 bg-primary text-primary-foreground shadow hover:bg-primary/90">
    <header className="sticky top-0 z-60 w-full bg-foreground text-primary-foreground shadow ">
      <div className="container flex h-16 items-center">
        <div className="mr-4 hidden md:flex">
          <div className="mr-2">
            <img
              src={logo}
              alt="PDF Logo"
              className="rounded-md"
              style={{ height: "35px", width: "35px" }}
            />
          </div>
          <div className="mr-6 flex items-center space-x-2">
            <span
              className="hidden font-bold sm:inline-block"
              style={{ fontSize: "20px" }}
            >
              Quickstore
            </span>
          </div>
          {/* </link> */}
          <nav className="flex items-center space-x-6 text-sm font-medium"></nav>
        </div>
      </div>
    </header>
  );
}

export default Header;
