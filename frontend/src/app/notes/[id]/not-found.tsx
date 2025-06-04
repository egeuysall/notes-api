import Link from "next/link";

export default function NotFound() {
  return (
    <main className="flex items-center justify-center h-screen">
      <section className="flex flex-col items-center justify-center gap-4 w-[90vw] rounded-lg bg-neutral-900 p-6 min-h-72 text-center">
        <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold">
          Note Not Found
        </h1>
        <p className="text-sm opacity-50">
          The note you&apos;re looking for doesn&apos;t exist or has expired.
        </p>
        <Link
          href="/"
          className="mt-2 px-4 py-2 rounded-lg bg-white text-black"
        >
          <button className="font-bold">Back to Home</button>
        </Link>
      </section>
    </main>
  );
}
