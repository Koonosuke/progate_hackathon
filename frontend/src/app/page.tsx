"use client";

import { useRouter } from "next/navigation";

const Page = () => {
  const router = useRouter();

  const handleClick = () => {
    router.push("/chat");
  };

  return (
    <div className="flex items-center justify-center h-screen">
      <button
        onClick={handleClick}
        className="bg-blue-500 hover:bg-blue-600 text-white px-6 py-3 rounded-lg shadow-md transition"
      >
        チャット画面へ
      </button>
    </div>
  );
};

export default Page;
