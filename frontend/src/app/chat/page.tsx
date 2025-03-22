"use client";

import { useState } from "react";

export default function ChatPage() {
  const [query, setQuery] = useState("");

  type Sticky = {
    content: string;
    category: string;
    color: string;
    created_at: string;
  };

  const [response, setResponse] = useState<Sticky[]>([]);
  const [advice, setAdvice] = useState("");
  const [loading, setLoading] = useState(false);

  const handleSearch = async () => {
    if (!query) return;
    setLoading(true);
    setResponse([]);
    setAdvice("");

    try {
      const res = await fetch(
        `http://localhost:8080/search_with_advice?query=${encodeURIComponent(
          query
        )}`
      );
      const data = await res.json();

      setResponse(data.results || []);
      setAdvice(data.advice || "");
    } catch (err) {
      console.error("Search failed:", err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="p-4 max-w-xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">ChatBot（振り返り検索）</h1>

      <div className="flex mb-4">
        <input
          className="flex-1 border p-2 rounded-l"
          placeholder="例: 会議長い、改善したい"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
        <button
          className="bg-blue-500 text-white px-4 rounded-r"
          onClick={handleSearch}
          disabled={loading}
        >
          {loading ? "検索中..." : "検索"}
        </button>
      </div>

      {advice && (
        <div className="bg-yellow-100 border-l-4 border-yellow-500 text-yellow-700 p-4 mb-4 rounded shadow">
          <p className="font-semibold">💡 アドバイス:</p>
          <p>{advice}</p>
        </div>
      )}

      <div className="space-y-2">
        {response.map((item, index) => (
          <div key={index} className="border rounded p-2 shadow">
            <p>
              <strong>内容:</strong> {item.content}
            </p>
            <p>
              <strong>カテゴリ:</strong> {item.category}
            </p>
            <p>
              <strong>作成日時:</strong> {item.created_at}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
}
