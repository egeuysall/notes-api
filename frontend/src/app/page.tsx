"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import type { FormEvent } from "react";

export default function Home() {
  const [noteId, setNoteId] = useState("");
  const router = useRouter();

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    router.push(`/notes/${encodeURIComponent(noteId)}`);
  };

  return (
    <main className="flex items-center justify-center h-screen">
      <section className="flex flex-col gap-6 w-[90vw]">
        <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold text-center">
          Notes API
        </h1>
        <form
          className="grid md:grid-cols-4 flex-col gap-4"
          onSubmit={handleSubmit}
        >
          <input
            className="border border-neutral-700 rounded-lg p-3 md:col-span-3"
            type="text"
            placeholder="Enter note id"
            value={noteId}
            onChange={(e) => setNoteId(e.target.value)}
          />
          <button
            className="bg-white text-black p-3 font-bold rounded-lg"
            type="submit"
          >
            Get Note
          </button>
        </form>
      </section>
    </main>
  );
}
