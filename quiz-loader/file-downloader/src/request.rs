use tokio::prelude::*;
// use hyper::body::HttpBody as _;
// use hyper::Client;
// use hyper_tls::HttpsConnector;

// pub async fn save_url_data_into_file(
//     url: &str,
//     file: &mut tokio::fs::File,
// ) -> Result<(), Box<dyn std::error::Error>> {
//     let https = HttpsConnector::new();
//     let client = Client::builder().build::<_, hyper::Body>(https);

//     let uri = url.parse::<hyper::Uri>()?;
//     let mut response = client.get(uri).await?;

//     while let Some(chunk) = response.data().await {
//         let data = chunk?;
//         file.write_all(data.as_ref()).await?;
//     }

//     Ok(())
// }

pub async fn save(
    url: &str,
    file: &mut tokio::fs::File,
) -> Result<(), Box<dyn std::error::Error>> {
    let client = reqwest::Client::new();
    let response = client.get(url).send().await?;

    if response.status() == reqwest::StatusCode::OK {
        let data = response.bytes().await?;
        file.write_all(data.as_ref()).await?;
    }

    Ok(())
}
