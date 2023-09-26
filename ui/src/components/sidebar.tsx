import Link from "next/link";
import { HomeIcon } from "@radix-ui/react-icons";

export function Sidebar() {
  return (
    <div className="w-64 flex flex-col px-6 py-8 space-y-4 border-r border-r-gray-800">
      <Link
        href="/"
        className="flex items-center px-4 py-1.5 rounded hover:bg-gray-800 duration-100"
      >
        <HomeIcon className="mr-2" />
        Home
      </Link>
    </div>
  );
}
