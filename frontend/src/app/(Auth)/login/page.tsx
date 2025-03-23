"use client";

import { useUser } from "@/app/context/UserContext"; // Contextをimport
import {
  AuthenticationDetails,
  CognitoUser,
  userPool,
} from "@/app/lib/cognito"; // あなたのcognito設定ファイル
import Link from "next/link";
import { useRouter } from "next/navigation"; // 任意で画面遷移したい場合に使用
import { useState } from "react";

export default function Login() {
  const [inputEmail, setInputEmail] = useState("");
  const [password, setPassword] = useState("");
  const { setAccessToken, setEmail } = useUser(); // Contextからsetterを取得
  const router = useRouter(); // 画面遷移したいとき用（任意）
  const clientId = "1cf6e5bu3c4asa0m5dfo5blkjq";
  const cognitoDomain =
    "https://us-west-2eyqu6wtem.auth.us-west-2.amazoncognito.com";
  const redirectUri = "http://localhost:3000/api/auth/callback/google"; // ← ✅スラッシュなしに戻す
  const loginUrl = `${cognitoDomain}/oauth2/authorize?identity_provider=Google&response_type=code&client_id=${clientId}&redirect_uri=${encodeURIComponent(
    redirectUri
  )}&scope=email+openid+profile`;

  console.log(loginUrl);

  const handleLogin = () => {
    const userData = {
      Username: inputEmail,
      Pool: userPool,
    };

    const authDetails = new AuthenticationDetails({
      Username: inputEmail,
      Password: password,
    });

    const cognitoUser = new CognitoUser(userData);

    cognitoUser.authenticateUser(authDetails, {
      onSuccess: (result) => {
        const token = result.getIdToken().getJwtToken();
        const claims = result.getIdToken().decodePayload();

        // ✅ Contextに格納
        setAccessToken(token);
        setEmail(claims.email);
        console.log(token);

        console.log("✅ Login success:", claims.email);
        alert(`ようこそ ${claims.email} さん！`);

        // 任意でページ遷移
        router.push("/"); // ← 任意
      },

      onFailure: (err) => {
        console.error("❌ Login failed:", err);
        alert("ログインに失敗しました");
      },
    });
  };

  return (
    <div className="h-screen flex justify-center items-center bg-gray-100">
      <div className="flex flex-col justify-center items-center p-8 bg-white rounded-lg shadow-md">
        <h2 className="text-2xl font-bold mb-6">ログイン</h2>
        <input
          type="email"
          placeholder="メールアドレス"
          value={inputEmail}
          onChange={(e) => setInputEmail(e.target.value)}
          className="mb-4 px-4 py-2 border border-gray-300 rounded w-64"
        />
        <input
          type="password"
          placeholder="パスワード"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="mb-6 px-4 py-2 border border-gray-300 rounded w-64"
        />
        <button
          onClick={handleLogin}
          className="bg-purple-400 hover:bg-purple-500 text-white px-6 py-2 rounded transition"
        >
          送信
        </button>
        <div>
          <Link href={"/signup"}>アカウントをお持ちでないですか？</Link>
        </div>
      </div>
    </div>
  );
}
