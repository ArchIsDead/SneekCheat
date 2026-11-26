const tt = new TextEncoder();
const kk = async (pw) =>
  crypto.subtle.importKey("raw", await crypto.subtle.digest("SHA-256", tt.encode(pw)), "AES-GCM", false, ["encrypt", "decrypt"]);

export default {
  async fetch(rq, ev) {
    if (rq.method !== "POST") return new Response("yo");
    const kx = await kk(ev.KEY || "loser123");
    let qq;
    try {
      const rr = Uint8Array.from(atob((await rq.text()).trim()), (cc) => cc.charCodeAt(0));
      const pp = await crypto.subtle.decrypt({ name: "AES-GCM", iv: rr.slice(0, 12) }, kx, rr.slice(12));
      qq = JSON.parse(new TextDecoder().decode(pp));
    } catch {
      return new Response("ga boleh", { status: 401 });
    }
    const bk = async (oo) => {
      const ii = crypto.getRandomValues(new Uint8Array(12));
      const zz = new Uint8Array(await crypto.subtle.encrypt({ name: "AES-GCM", iv: ii }, kx, tt.encode(JSON.stringify(oo))));
      let ss = String.fromCharCode(...ii);
      for (const xx of zz) ss += String.fromCharCode(xx);
      return new Response(btoa(ss));
    };
    if (!/^https:\/\//.test(qq.url)) return bk({ ok: false, text: "url-nya harus https" });
    try {
      const rs = await fetch(qq.url, {
        method: qq.method,
        headers: qq.headers,
        body: qq.method === "GET" ? undefined : qq.body,
      });
      const tx = await rs.text();
      let an;
      try {
        an = qq.path.split(".").reduce((aa, kv) => aa?.[kv], JSON.parse(tx));
      } catch {}
      if (typeof an !== "string") return bk({ ok: false, text: `http ${rs.status}: ${tx.slice(0, 400)}` });
      return bk({ ok: true, text: an });
    } catch (ee) {
      return bk({ ok: false, text: String(ee) });
    }
  },
};
