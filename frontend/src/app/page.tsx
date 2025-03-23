"use client";
import { isSameDayJST, toDateStringJST } from "@/app/utils/date";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { Calendar } from "./components/Calender";
import { DailyReport } from "./components/DailyReport";
import { useUser } from "./context/UserContext";
import { StickyNote } from "./types/report";

export default function Home() {
  const { accessToken } = useUser();
  //const [message, setMessage] = useState("");
  const [selectedDate, setSelectedDate] = useState(new Date());
  //const [reports, setReports] = useState<StickyNote[]>([]);
  const [improvePoints, setImprovePoints] = useState<string[]>([]);
  const [repentedPoints, setRepentedPoints] = useState<string[]>([]);

  const router = useRouter();

  {
    /*const toUTCDateString = (date: string | Date) =>
    format(new Date(date), "yyyy-MM-dd");*/
  }

  const handleDateSelect = async (date: Date) => {
    setSelectedDate(date);

    try {
      const res = await fetch("http://localhost:8080/stickies", {
        headers: {
          Authorization: `Bearer ${accessToken}`,
        },
      });

      if (!res.ok) {
        console.error("❌ /stickies 取得失敗:", res.status);
        return;
      }

      const data: StickyNote[] = await res.json();

      const selectedDateStr = toDateStringJST(date); // JSTで文字列に変換
      console.log("フォーマットされたでーた" + selectedDateStr);

      // 選択日と一致する付箋のみ抽出
      const filtered = data.filter((item) =>
        isSameDayJST(item.created_at, date)
      );
      console.log("filterd" + filtered);

      const improvements = filtered
        .filter((item) => item.category === "改善点")
        .map((item) => item.content);

      const repenteds = filtered
        .filter((item) => item.category === "反省点")
        .map((item) => item.content);
      setImprovePoints(improvements);
      setRepentedPoints(repenteds);
      console.log(data);
      console.log(repentedPoints);
      console.log(improvePoints);
      console.log("選択された日付 JST:", selectedDateStr);
      filtered.forEach((item) => {
        console.log(
          "created_at:",
          item.created_at,
          "→ JST:",
          toDateStringJST(item.created_at)
        );
      });
    } catch (err) {
      if (err instanceof Error) {
        console.error("付箋取得エラー:", err.message);
      } else {
        console.error("付箋取得エラー:", err);
      }
    }
  };

  {
    /*const getReport = (date: Date): DailyReportType | undefined => {
    const selectedDateStr = toUTCDateString(date); // UTCのまま日付を抽出

    const reflectionNotes = reports.filter(
      (r) =>
        toUTCDateString(r.created_at) === selectedDateStr &&
        r.category === "反省点"
    );
    const improvementNotes = reports.filter(
      (r) =>
        toUTCDateString(r.created_at) === selectedDateStr &&
        r.category === "改善点"
    );

    const reflection = reflectionNotes.map((r) => r.content).join("\n");
    const improvement = improvementNotes.map((r) => r.content).join("\n");

    if (!reflection && !improvement) return undefined;

    return {
      date: selectedDateStr,
      reflection,
      improvement,
    };
  };*/
  }

  {
    /*useEffect(() => {
    if (isLoading) return;

    if (!accessToken) {
      router.push("/login");
      return;
    }

    // ユーザー取得
    fetch("http://localhost:8080/api/user")
      .then(async (res) => {
        if (!res.ok) {
          const text = await res.text();
          if (res.status === 401 || text.includes("Invalid token")) {
            console.warn(
              "アクセストークンが無効です。ログイン画面へ遷移します。"
            );
            router.push("/login");
            return;
          }
          throw new Error(text);
        }
        return res.json();
      })
      .then((data) => {
        if (data) setMessage(data.message);
      })
      .catch((err) => console.error("Fetch error:", err.message));

    // 付箋データ取得
    fetch("http://localhost:8080/stickies", {
      headers: {
        Authorization: `Bearer ${accessToken}`,
      },
    })
      .then((res) => {
        if (!res.ok) {
          console.error("❌ /stickies 取得失敗:", res.status);
          return [];
        }
        return res.json();
      })
      .then((data) => {
        console.log("✅ /stickies 取得成功:", data);

        setReports(data);
        console.log(reports);
      })
      .catch((err) => console.error("付箋取得エラー:", err));
  }, [accessToken, isLoading, router]);
  */
  }

  return (
    <div className="min-h-screen bg-gray-100">
      <header className="bg-white border-b">
        <div className="max-w-7xl mx-auto px-4 py-4 flex items-center justify-between">
          <h1 className="text-xl font-semibold">ホーム</h1>
          <div className="flex items-center gap-4">
            <h2 className="text-lg text-gray-700">こんにちは</h2>
            <button
              onClick={() => router.push("/chat")}
              className="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition"
            >
              チャットへ
            </button>
          </div>
        </div>
      </header>

      <main className="max-w-7xl mx-auto px-4 py-8 h-[calc(100vh-73px)]">
        <div className="flex gap-8 h-full">
          <div className="w-1/2">
            <Calendar
              selectedDate={selectedDate}
              onDateSelect={handleDateSelect}
            />
          </div>
          <div className="w-1/2">
            <DailyReport
              date={selectedDate}
              improvePoints={improvePoints}
              repentedPoints={repentedPoints}
            />
          </div>
        </div>
      </main>
    </div>
  );
}
