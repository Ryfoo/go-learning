# go-learning

Personal repository for tracking my Golang learning journey — from the very basics all the way to building real projects.

---

## 🔴 Repository Structure

```
go-learning/
├── 01_basics/
├── 02_advanced/
├── 03_standard_library/
├── 04_projects/
├── 05_notes/
└── 06_resources/
```

### 🟢 01_basics
Core Go fundamentals — syntax, variables, types, control structures, and functions.

### 🟢 02_advanced
Advanced concepts — structs, methods, interfaces, error handling, generics, and concurrency.

### 🟢 03_standard_library
Experiments and examples with Go's standard library (`fmt`, `os`, `net/http`, `time`, and more).

### 🟢 04_projects
Mini projects applying what I have learned — calculators, TODO apps, HTTP servers, and beyond.

### 🟢 05_notes
Personal notes, diagrams, and tips for deeper understanding.

### 🟢 06_resources
Books, courses, blogs, and other references guiding the learning process.

---

## 🟡 Current State

Diving into the language's fundamental features.

## 🟡 Goals

- Build a strong foundation in Go.
- Practice writing clean, idiomatic Go code.
- Apply concepts in small, runnable projects.
- Track progress and retain key insights efficiently.

---

## 🔴 Usage

- Browse folders to explore topics by category.
- Run examples in each folder to practice hands-on.
- Use `05_notes/` to review key concepts and tips.

---

## 🟢 Frontend Setup — React + TypeScript + Tailwind CSS

For any project in `04_projects/` that includes a frontend interface, follow this setup to scaffold a React app with TypeScript and Tailwind CSS v3.3.3.

---

### Prerequisites

Make sure you have the following installed before starting:

- [Node.js](https://nodejs.org/) (v18 or later recommended)
- npm (comes with Node) or [pnpm](https://pnpm.io/) / [yarn](https://yarnpkg.com/) if you prefer

---

### Step 1 — Scaffold the React + TypeScript app

Use Vite to create the project (fast, minimal, and the current standard):

```bash
npm create vite@latest my-app -- --template react-ts
cd my-app
npm install
```

> 🟡 Replace `my-app` with your actual project name.

---

### Step 2 — Install Tailwind CSS v3.3.3

Pin the exact version to stay on the LTS-stable v3 release:

```bash
npm install -D tailwindcss@3.3.3 postcss autoprefixer
npx tailwindcss init -p
```

The `-p` flag generates both `tailwind.config.js` and `postcss.config.js` in one step.

---

### Step 3 — Configure Tailwind to scan your source files

Open `tailwind.config.js` and update the `content` array so Tailwind knows which files to scan for class names:

```js
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
```

---

### Step 4 — Add Tailwind directives to your CSS

Replace the contents of `src/index.css` with the three Tailwind directives:

```css
@tailwind base;
@tailwind components;
@tailwind utilities;
```

---

### Step 5 — Make sure the CSS is imported

Open `src/main.tsx` and confirm that `index.css` is imported at the top (Vite includes this by default, but worth checking):

```tsx
import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
```

---

### Step 6 — Verify the setup

Replace the contents of `src/App.tsx` with a quick smoke test:

```tsx
function App() {
  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <h1 className="text-4xl font-bold text-blue-600">
        React + TypeScript + Tailwind CSS
      </h1>
    </div>
  )
}

export default App
```

Then start the dev server:

```bash
npm run dev
```

Open [http://localhost:5173](http://localhost:5173) — if you see the styled heading, everything is working.

---

### Final Project Structure

After setup, your frontend project will look like this:

```
my-app/
├── public/
├── src/
│   ├── assets/
│   ├── App.tsx
│   ├── main.tsx
│   └── index.css        ← Tailwind directives live here
├── index.html
├── tailwind.config.js   ← Tailwind configuration
├── postcss.config.js    ← PostCSS configuration (auto-generated)
├── tsconfig.json
└── vite.config.ts
```

---

### Quick Reference — Useful Commands

| Command | What it does |
|---------|-------------|
| `npm run dev` | Start the development server |
| `npm run build` | Build for production |
| `npm run preview` | Preview the production build locally |
| `npx tailwindcss --help` | Check Tailwind CLI options |

---

> 🟢 You are all set. From here, build out your components in `src/`, use Tailwind utility classes freely, and connect to your Go backend via `fetch` or a library like [axios](https://axios-http.com/).
