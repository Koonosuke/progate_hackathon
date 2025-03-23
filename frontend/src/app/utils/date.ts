// src/app/utils/date.ts

// JSTで "yyyy-MM-dd" を取り出す（9時間ずらすだけ）
export const toDateStringJST = (date: string | Date): string => {
  const jst = new Date(new Date(date).getTime() + 9 * 60 * 60 * 1000);
  return jst.toISOString().split("T")[0]; // yyyy-MM-dd だけ取り出す
};

// 2つの日付が JST 日付で一致しているか比較
export const isSameDayJST = (a: Date | string, b: Date | string): boolean => {
  return toDateStringJST(a) === toDateStringJST(b);
};
