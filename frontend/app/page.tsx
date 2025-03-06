import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <div className="grid grid-rows-[20px_1fr_20px] justify-center items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-gist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <div className="w-60 h-60 rounded-lg shadow-2xl overflow-hidden">
          <Image
            src="/cook.png" // Update with your actual image path
            alt="Profile"
            width={240}
            height={240}
            className="object-cover w-full h-full"
          />
        </div>
        <div className="flex gap-4 items-center flex-col sm:flex-row">
        </div>
      </main>
    </div>
  );
}