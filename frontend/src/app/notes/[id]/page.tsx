import React from "react";
import { notFound } from "next/navigation";

export default async function Home({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  let note: { id: string; note: string } | null = null;

  try {
    const response = await fetch(
      `http://localhost:8080/v1/notes/${encodeURIComponent(id)}`,
      {
        method: "GET",
        headers: {
          Authorization: `Bearer ${process.env.API_KEY}`,
          "Content-Type": "application/json",
        },
      }
    );

    if (!response.ok) {
      throw new Error(`HTTP error. Status code: ${response.status}`);
    }

    note = await response.json();
  } catch (error) {
    console.error(`Failed to fetch note: ${error}`);
    notFound();
  }

  if (!note) {
    notFound();
  }

  return (
    <main className="flex items-center justify-center h-screen bg-neutral-950">
      <section className="flex flex-col gap-3 w-[90vw] rounded-lg bg-neutral-900 p-6 min-h-72 items-center justify-center">
        <section className="flex flex-col gap-1">
          <h1 className="text-2xl md:text-3xl lg:text-4xl font-bold text-center">
            Note Details
          </h1>
          <p className="text-center text-sm opacity-50">
            <strong>ID:</strong> {note.id}
          </p>
        </section>
        <p className="text-center line-clamp-8 break-words">{note.note}</p>
      </section>
    </main>
  );
}
