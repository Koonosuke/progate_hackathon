import { DailyReport as DailyReportType } from "@/app/types/report";
import { format } from "date-fns";
import { ja } from "date-fns/locale";

interface DailyReportProps {
  date: Date;
  report?: DailyReportType;
  improvePoints: string[];
  repentedPoints: string[];
}
{
  /* 
// HTMLタグ除去のユーティリティ
const stripHtml = (html: string) => {
  return html.replace(/<[^>]*>?/gm, "").trim();
};
*/
}

// 改行と空行除去
{
  /*const formatContent = (text: string): string[] => {
  return stripHtml(text)
    .split("\n")
    .map((line) => line.trim())
    .filter((line) => line.length > 0);
};*/
}

export function DailyReport({
  date,
  repentedPoints,
  improvePoints,
}: DailyReportProps) {
  return (
    <div className="w-full h-full bg-white rounded-lg p-8">
      <h2 className="text-2xl font-semibold mb-8">
        {format(date, "M月d日の日報の振り返り", { locale: ja })}
      </h2>

      <div className="space-y-8">
        <section>
          <h3 className="text-lg font-semibold mb-4">反省点</h3>
          {repentedPoints.length > 0 ? (
            <ul className="list-disc list-inside text-gray-700 space-y-1">
              {repentedPoints.map((item, idx: number) => (
                <li key={idx}>{item.replace(/^反省点/, "").trim()}</li>
              ))}
            </ul>
          ) : (
            <p className="text-gray-700">
              この日に記録された反省点はありません。
            </p>
          )}
        </section>

        <section>
          <h3 className="text-lg font-semibold mb-4">改善点</h3>
          {improvePoints.length > 0 ? (
            <ul className="list-disc list-inside text-gray-700 space-y-1">
              {improvePoints.map((item, idx) => (
                <li key={idx}>{item.replace(/^改善点/, "").trim()}</li>
              ))}
            </ul>
          ) : (
            <p className="text-gray-700">
              この日に記録された改善点はありません。
            </p>
          )}
        </section>
      </div>
    </div>
  );
}
