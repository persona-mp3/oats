package mp3.persona.oats.utils;


import org.springframework.util.MultiValueMap;
import org.springframework.web.util.UriComponentsBuilder;

import java.net.*;

public class Utils {
	public static String parseUrl(String url) throws Exception {
		String toString = "";
		try {
			URI _url = new URI(url);
			toString = _url.toString();
		} catch (URISyntaxException err) {
			System.err.printf("error in parsing url", err);
		} catch (Exception err) {
			throw err;
		}

		return toString;
	}

	public static String parseUrl(String url, MultiValueMap<String, String> params) throws Exception {
		String toString = "";
		try {
			URI _uri = UriComponentsBuilder
					.fromUriString(parseUrl(url))
					.queryParams(params)
					.build().toUri();

			toString = _uri.toString();

		} catch (URISyntaxException err) {
			System.err.printf("error in parsing url", err);
		} catch (Exception err) {
			throw err;
		}

		return toString;

	}
}
