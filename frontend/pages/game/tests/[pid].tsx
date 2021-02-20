import { GetStaticPaths, GetStaticProps } from "next"
import { Dispatch, SetStateAction, useState } from "react"
import styled from "styled-components"
import MessageModal from "../../../components/MessageModal"

type MediaInfo = {
    id: string;
    uri: string;
}
import { useRouter } from 'next/router'

const Post = () => {
  const router = useRouter()
  const { pid } = router.query

  return <p>Post: {pid}</p>
}

export default Post
