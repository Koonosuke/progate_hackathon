'use client';

import React, { useState } from 'react';
import {
  CognitoUserAttribute,
  userPool,
  CognitoUser,
} from '@/app/lib/cognito';
import { useRouter } from 'next/navigation'; // ← 追加
import Link from 'next/link';

export default function SignUp() {
  const router = useRouter(); // ← 追加
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmationCode, setConfirmationCode] = useState('');
  const [step, setStep] = useState<'signup' | 'confirm'>('signup');

  const handleSignUp = () => {
    const attributeList = [
      new CognitoUserAttribute({
        Name: 'email',
        Value: email,
      }),
    ];

    userPool.signUp(email, password, attributeList, [], (err, result) => {
      if (err) {
        console.error('SignUp error:', err);
        alert(err.message || 'サインアップ失敗');
        return;
      }

      console.log('SignUp success:', result);
      setStep('confirm');
    });
  };

  const handleConfirm = () => {
    const user = new CognitoUser({
      Username: email,
      Pool: userPool,
    });

    user.confirmRegistration(confirmationCode, true, (err, result) => {
      if (err) {
        console.error('Confirm error:', err);
        alert(err.message || '確認失敗');
        return;
      }

      console.log('Confirm success:', result);
      alert('登録完了！ログインしてください。');
      router.push("/login"); // ← ログイン画面に遷移
    });
  };

  return (
    <div className="h-screen flex justify-center items-center bg-gray-100 dark:bg-gray-900 transition-colors">
      <div className="bg-white dark:bg-gray-800 text-black dark:text-white rounded-lg shadow-md p-8 w-80">
        {step === 'signup' ? (
          <>
            <h2 className="text-2xl font-bold mb-6 text-center">サインアップ</h2>
            <input
              type="email"
              placeholder="メールアドレス"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="mb-4 px-4 py-2 border border-gray-300 dark:border-gray-600 rounded w-full bg-white dark:bg-gray-700 text-black dark:text-white"
            />
            <input
              type="password"
              placeholder="パスワード（8文字以上）"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mb-6 px-4 py-2 border border-gray-300 dark:border-gray-600 rounded w-full bg-white dark:bg-gray-700 text-black dark:text-white"
            />
            <button
              onClick={handleSignUp}
              className="bg-purple-400 hover:bg-purple-500 text-white px-4 py-2 rounded w-full transition"
            >
              登録
            </button>
            <div className='border-t-1 text-slate-300 p-1 m-1'>
              <Link href={"/login"} className='text-slate-500 hover:text-slate-900'>すでにアカウントをお持ちの場合</Link>
            </div>
          </>
        ) : (
          <>
            <h2 className="text-2xl font-bold mb-4 text-center">確認コードを入力</h2>
            <p className="text-sm mb-4 text-center text-gray-600 dark:text-gray-300">
              メールに届いた確認コードを入力してください。
            </p>
            <input
              type="text"
              placeholder="確認コード"
              value={confirmationCode}
              onChange={(e) => setConfirmationCode(e.target.value)}
              className="mb-6 px-4 py-2 border border-gray-300 dark:border-gray-600 rounded w-full bg-white dark:bg-gray-700 text-black dark:text-white"
            />
            <button
              onClick={handleConfirm}
              className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded w-full transition"
            >
              確認
            </button>
          </>
        )}
      </div>
    </div>
  );
}
