"use client";

import { useEffect } from "react";

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  useEffect(() => {
    console.error("Page error:", error);
  }, [error]);

  return (
    <main className="flex items-center justify-center h-screen">
      <section className="flex flex-col items-center justify-center gap-4 w-[90vw] rounded-lg bg-neutral-900 p-6 min-h-72 text-center">
        <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold">
          Something went wrong
        </h1>
        <p className="text-sm opacity-50">
          {error.message || "An unexpected error occurred."}
        </p>
        <button
          onClick={reset}
          className="px-4 py-2 bg-white text-black rounded-lg font-bold"
        >
          Try again
        </button>
      </section>
    </main>
  );
}
