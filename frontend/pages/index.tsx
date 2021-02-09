import Head from 'next/head'
import styles from '../styles/Home.module.css'

export default function Home() {
  return (
    <div className={styles.container}>
      <Head>
        <title>Learnify - Learn by playing!</title>
        <link rel="icon" href="favicon.ico" />
        <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Roboto:300,400,500,700&display=swap" />
        <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons" />
        <meta
          name="viewport"
          content="minimum-scale=1, initial-scale=1, width=device-width"
        />
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
