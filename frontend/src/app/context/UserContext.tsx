"use client";

import { getCurrentUser } from "@/app/lib/cognito";
import { CognitoUserSession } from "amazon-cognito-identity-js";
import React, { createContext, useContext, useEffect, useState } from "react";

type UserContextType = {
  accessToken: string | null;
  email: string | null;
  isLoading: boolean;
  setAccessToken: (token: string) => void;
  setEmail: (email: string) => void;
};

const UserContext = createContext<UserContextType>({
  accessToken: null,
  email: null,
  isLoading: true,
  setAccessToken: () => {},
  setEmail: () => {},
});

export const UserProvider = ({ children }: { children: React.ReactNode }) => {
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [email, setEmail] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const user = getCurrentUser();
    if (!user) {
      setIsLoading(false); // ユーザーがいない場合も読み込み完了
      return;
    }

    user.getSession((err: Error, session: CognitoUserSession | null) => {
      if (err || !session?.isValid()) {
        console.warn("セッション無効または取得失敗:", err);
        setIsLoading(false); // エラー時も完了とみなす
        return;
      }

      setAccessToken(session.getAccessToken().getJwtToken());

      const claims = session.getIdToken().decodePayload();
      setEmail(claims.email);

      setIsLoading(false); // 読み込み完了
    });
  }, []);

  return (
    <UserContext.Provider
      value={{ accessToken, email, isLoading, setAccessToken, setEmail }}
    >
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
