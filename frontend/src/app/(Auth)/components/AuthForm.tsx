"use client";
import "./AuthForm.css";

export default function UserAuthForm() {
  return (
    <div className="login-form-container">
      <form className="login-form" action="">
        <div className="form-fields">
          <div className="form-group">
            <label htmlFor="email">メールアドレス</label>
            <input
              id="email"
              placeholder="name@example.com"
              type="email"
              className="form-input"
            />
            <label htmlFor="password">パスワード</label>
            <input
              id="password"
              placeholder="ハスワードを入れてください"
              type="password"
              className="form-input"
            />
          </div>
          <button className="login-button">メールアドレスでログイン</button>
        </div>
      </form>

      <div className="divider-container">
        <div className="divider-line-wrapper">
          <span className="divider-line"></span>
        </div>
        <div className="divider-text-wrapper">
          <span className="divider-text">または</span>
        </div>
      </div>

      {/* OAuth認証やりたいときに使うかも
            <div className="flex flex-col gap-3">
            <button 
            className={cn(buttonVariants({variant: "outline"}))} 
            onClick={() => {
                setIsGithubLoading(true)
                signIn("github")
            }}
            >

            {!isGithubLoading ? <Icon.github/>: <Icon.spinner className="animate-spin"/>}
            Github</button> 
            <button 
            className={cn(buttonVariants({variant: "outline"}))} 
            onClick={() => {
                setIsGoogleLoading(true)
                signIn("google")
            }}
            >
            {!isGoogleLoading ? <Icon.google/>: <Icon.spinner className="animate-spin"/>}
            Google</button> 
            </div>
            */}
    </div>
  );
}
