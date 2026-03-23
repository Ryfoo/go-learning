import { useState, useEffect } from "react";

const REPO = "Ryfoo/go-learning";
const BRANCH = "main";

const FOLDERS: { id: string; label: string }[] = [
  { id: "basics", label: "Basics" },
  { id: "advanced_topics", label: "Advanced Topics" },
  { id: "stdlib", label: "Stdlib" },
  { id: "project", label: "Project" },
  { id: "notes", label: "Notes" },
];

const RAW = (folder: string): string =>
  `https://raw.githubusercontent.com/${REPO}/${BRANCH}/${folder}/README.md`;

function MarkdownRenderer({ content }: { content: string }) {
  const render = (text: string): string => {
    if (!text) return "";
    return text
      .replace(/^#{6}\s(.+)$/gm, '<h6 class="md-h6">$1</h6>')
      .replace(/^#{5}\s(.+)$/gm, '<h5 class="md-h5">$1</h5>')
      .replace(/^#{4}\s(.+)$/gm, '<h4 class="md-h4">$1</h4>')
      .replace(/^#{3}\s(.+)$/gm, '<h3 class="md-h3">$1</h3>')
      .replace(/^#{2}\s(.+)$/gm, '<h2 class="md-h2">$1</h2>')
      .replace(/^#{1}\s(.+)$/gm, '<h1 class="md-h1">$1</h1>')
      .replace(/```(\w+)?\n([\s\S]*?)```/gm, '<pre class="md-pre"><code>$2</code></pre>')
      .replace(/`([^`]+)`/g, '<code class="md-code">$1</code>')
      .replace(/\*\*(.+?)\*\*/g, "<strong>$1</strong>")
      .replace(/\*(.+?)\*/g, "<em>$1</em>")
      .replace(/^\s*[-*]\s(.+)$/gm, '<li class="md-li">$1</li>')
      .replace(/(<li[\s\S]*?<\/li>)/g, '<ul class="md-ul">$1</ul>')
      .replace(/\[([^\]]+)\]\(([^)]+)\)/g, '<a class="md-a" href="$2" target="_blank" rel="noreferrer">$1</a>')
      .replace(/^(?!<[hupla]|<\/[hupla]|<pre|<\/pre)(.+)$/gm, '<p class="md-p">$1</p>')
      .replace(/\n{2,}/g, "");
  };

  return (
    <div
      className="md-content"
      dangerouslySetInnerHTML={{ __html: render(content) }}
    />
  );
}

export default function App() {
  const [active, setActive] = useState<string>("basics");
  const [contents, setContents] = useState<Record<string, string>>({});
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (contents[active]) return;
    setLoading(true);
    setError(null);
    fetch(RAW(active))
      .then((r) => {
        if (!r.ok) throw new Error(`No README found in /${active}`);
        return r.text();
      })
      .then((text) => {
        setContents((prev) => ({ ...prev, [active]: text }));
        setLoading(false);
      })
      .catch((e: Error) => {
        setError(e.message);
        setLoading(false);
      });
  }, [active]);

  const current = FOLDERS.find((f) => f.id === active);

  return (
    <div style={styles.root}>
      <style>{css}</style>

      {/* Sidebar */}
      <aside style={styles.sidebar}>
        <div style={styles.brand}>
          <span style={styles.brandGo}>Go</span>
          <span style={styles.brandText}>Learning</span>
        </div>

        <div style={styles.repoTag}>
          <span style={styles.repoIcon}>⌥</span>
          <span style={styles.repoLabel}>{REPO}</span>
        </div>

        <nav style={styles.nav}>
          <p style={styles.navLabel}>sections</p>
          {FOLDERS.map((f) => (
            <button
              key={f.id}
              onClick={() => setActive(f.id)}
              style={{
                ...styles.navItem,
                ...(active === f.id ? styles.navItemActive : {}),
              }}
            >
              <span style={styles.navIcon}>
                {active === f.id ? "◆" : "◇"}
              </span>
              {f.label}
            </button>
          ))}
        </nav>

        <div style={styles.sidebarFooter}>
          <span style={styles.footerDot} />
          <span style={styles.footerText}>github / {BRANCH}</span>
        </div>
      </aside>

      {/* Main */}
      <main style={styles.main}>
        <header style={styles.header}>
          <div style={styles.breadcrumb}>
            <span style={styles.breadcrumbRepo}>{REPO}</span>
            <span style={styles.breadcrumbSep}>/</span>
            <span style={styles.breadcrumbActive}>{current?.id}</span>
            <span style={styles.breadcrumbSep}>/</span>
            <span style={styles.breadcrumbFile}>README.md</span>
          </div>
        </header>

        <div style={styles.content}>
          {loading && (
            <div style={styles.state}>
              <div style={styles.spinner} className="spin" />
              <span style={styles.stateText}>fetching README...</span>
            </div>
          )}

          {error && !loading && (
            <div style={styles.state}>
              <span style={styles.errorIcon}>✕</span>
              <span style={styles.errorText}>{error}</span>
            </div>
          )}

          {!loading && !error && contents[active] && (
            <MarkdownRenderer content={contents[active]} />
          )}
        </div>
      </main>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  root: {
    display: "flex",
    height: "100vh",
    background: "#0d1117",
    color: "#c9d1d9",
    fontFamily: "'JetBrains Mono', 'Fira Code', monospace",
    overflow: "hidden",
  },
  sidebar: {
    width: "220px",
    minWidth: "220px",
    background: "#010409",
    borderRight: "1px solid #21262d",
    display: "flex",
    flexDirection: "column",
    padding: "24px 0",
  },
  brand: {
    display: "flex",
    alignItems: "center",
    gap: "8px",
    padding: "0 20px 24px",
    borderBottom: "1px solid #21262d",
  },
  brandGo: {
    fontSize: "13px",
    fontWeight: "700",
    background: "#00acd7",
    color: "#0d1117",
    padding: "2px 7px",
    borderRadius: "4px",
    letterSpacing: "0.5px",
  },
  brandText: {
    fontSize: "13px",
    color: "#8b949e",
    fontWeight: "400",
    letterSpacing: "0.3px",
  },
  repoTag: {
    display: "flex",
    alignItems: "center",
    gap: "6px",
    padding: "14px 20px",
    borderBottom: "1px solid #21262d",
  },
  repoIcon: {
    fontSize: "11px",
    color: "#00acd7",
  },
  repoLabel: {
    fontSize: "11px",
    color: "#8b949e",
    letterSpacing: "0.3px",
  },
  nav: {
    flex: 1,
    padding: "16px 0",
  },
  navLabel: {
    fontSize: "10px",
    color: "#484f58",
    letterSpacing: "1.2px",
    textTransform: "uppercase",
    padding: "0 20px",
    margin: "0 0 8px",
  },
  navItem: {
    display: "flex",
    alignItems: "center",
    gap: "10px",
    width: "100%",
    padding: "8px 20px",
    background: "none",
    border: "none",
    color: "#8b949e",
    fontSize: "12px",
    cursor: "pointer",
    textAlign: "left",
    letterSpacing: "0.2px",
    transition: "all 0.15s",
  },
  navItemActive: {
    color: "#e6edf3",
    background: "#161b22",
    borderRight: "2px solid #00acd7",
  },
  navIcon: {
    fontSize: "8px",
    color: "#00acd7",
    minWidth: "10px",
  },
  sidebarFooter: {
    display: "flex",
    alignItems: "center",
    gap: "6px",
    padding: "16px 20px 0",
    borderTop: "1px solid #21262d",
  },
  footerDot: {
    width: "6px",
    height: "6px",
    borderRadius: "50%",
    background: "#3fb950",
    display: "inline-block",
  },
  footerText: {
    fontSize: "10px",
    color: "#484f58",
    letterSpacing: "0.3px",
  },
  header: {
    padding: "16px 32px",
    borderBottom: "1px solid #21262d",
    background: "#0d1117",
  },
  breadcrumb: {
    display: "flex",
    alignItems: "center",
    gap: "6px",
    fontSize: "12px",
  },
  breadcrumbRepo: { color: "#58a6ff" },
  breadcrumbSep: { color: "#484f58" },
  breadcrumbActive: { color: "#8b949e" },
  breadcrumbFile: { color: "#8b949e" },
  main: {
    flex: 1,
    display: "flex",
    flexDirection: "column",
    overflow: "hidden",
    background: "#0d1117",
  },
  content: {
    flex: 1,
    overflowY: "auto",
    padding: "32px 48px",
    maxWidth: "820px",
  },
  state: {
    display: "flex",
    alignItems: "center",
    gap: "12px",
    padding: "48px 0",
    color: "#8b949e",
  },
  stateText: {
    fontSize: "13px",
    color: "#484f58",
  },
  spinner: {
    width: "16px",
    height: "16px",
    border: "2px solid #21262d",
    borderTop: "2px solid #00acd7",
    borderRadius: "50%",
  },
  errorIcon: {
    fontSize: "14px",
    color: "#f85149",
  },
  errorText: {
    fontSize: "13px",
    color: "#8b949e",
  },
};

const css = `
  @import url('https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;500;700&display=swap');

  * { box-sizing: border-box; margin: 0; padding: 0; }

  ::-webkit-scrollbar { width: 6px; }
  ::-webkit-scrollbar-track { background: #010409; }
  ::-webkit-scrollbar-thumb { background: #21262d; border-radius: 3px; }

  @keyframes spin { to { transform: rotate(360deg); } }
  .spin { animation: spin 0.8s linear infinite; }

  .md-content { color: #c9d1d9; line-height: 1.75; font-family: 'JetBrains Mono', monospace; }
  .md-h1 { font-size: 22px; font-weight: 700; color: #e6edf3; margin: 0 0 20px; padding-bottom: 10px; border-bottom: 1px solid #21262d; }
  .md-h2 { font-size: 17px; font-weight: 600; color: #e6edf3; margin: 28px 0 12px; padding-bottom: 6px; border-bottom: 1px solid #21262d; }
  .md-h3 { font-size: 14px; font-weight: 600; color: #c9d1d9; margin: 20px 0 8px; }
  .md-h4, .md-h5, .md-h6 { font-size: 13px; font-weight: 600; color: #8b949e; margin: 16px 0 6px; }
  .md-p { font-size: 13px; color: #8b949e; margin: 0 0 12px; line-height: 1.8; }
  .md-pre { background: #161b22; border: 1px solid #21262d; border-radius: 6px; padding: 16px 20px; margin: 16px 0; overflow-x: auto; }
  .md-pre code { font-size: 12px; color: #e6edf3; font-family: 'JetBrains Mono', monospace; line-height: 1.7; white-space: pre; }
  .md-code { background: #161b22; color: #00acd7; font-size: 12px; padding: 2px 6px; border-radius: 4px; border: 1px solid #21262d; font-family: 'JetBrains Mono', monospace; }
  .md-ul { padding-left: 0; margin: 8px 0 12px; list-style: none; }
  .md-li { font-size: 13px; color: #8b949e; padding: 3px 0 3px 16px; position: relative; line-height: 1.7; }
  .md-li::before { content: '—'; position: absolute; left: 0; color: #00acd7; font-size: 11px; }
  .md-a { color: #58a6ff; text-decoration: none; font-size: 13px; }
  .md-a:hover { text-decoration: underline; }
  strong { color: #e6edf3; font-weight: 600; }
  em { color: #8b949e; font-style: italic; }
`;
