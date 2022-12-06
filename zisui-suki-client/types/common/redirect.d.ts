type Redirect =
| {
    statusCode: 301 | 302 | 303 | 307 | 308; // ステータスコード
    destination: string; // リダイレクト先のURL
    basePath?: false; // `basePath`を無効にします
  }
| {
    permanent: boolean; // 永続的なリダイレクト化のフラグ
    destination: string; // 同上
    basePath?: false; // 同上
  };