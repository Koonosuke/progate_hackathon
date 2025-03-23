export type DailyReport = {
  date: string;
  reflection: string; // ⬅ 文字列の配列に変更
  improvement: string;
};

export type StickyNote = {
  id: number;
  date?: string;
  category: "反省点" | "改善点";
  content: string;
  created_at: string; // ← Date型ではなく string にするのがポイント
};
