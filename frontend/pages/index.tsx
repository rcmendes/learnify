import Head from 'next/head'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Learnify - Learn by playing!</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        Main/Home
      </main>

      <footer className={styles.footer}>
        <a
          href="https://ciriguela.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          Ciriguela
        </a>
      </footer>
    </div>
  )
}
