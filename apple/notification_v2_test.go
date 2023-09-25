package apple

import (
	"encoding/json"
	"testing"

	"github.com/Adamxu0120/gopay/pkg/xlog"
)

var notificationV2Body = `{"signedPayload":"eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlFTURDQ0E3YWdBd0lCQWdJUWFQb1BsZHZwU29FSDBsQnJqRFB2OWpBS0JnZ3Foa2pPUFFRREF6QjFNVVF3UWdZRFZRUURERHRCY0hCc1pTQlhiM0pzWkhkcFpHVWdSR1YyWld4dmNHVnlJRkpsYkdGMGFXOXVjeUJEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURUxNQWtHQTFVRUN3d0NSell4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRJeE1EZ3lOVEF5TlRBek5Gb1hEVEl6TURreU5EQXlOVEF6TTFvd2daSXhRREErQmdOVkJBTU1OMUJ5YjJRZ1JVTkRJRTFoWXlCQmNIQWdVM1J2Y21VZ1lXNWtJR2xVZFc1bGN5QlRkRzl5WlNCU1pXTmxhWEIwSUZOcFoyNXBibWN4TERBcUJnTlZCQXNNSTBGd2NHeGxJRmR2Y214a2QybGtaU0JFWlhabGJHOXdaWElnVW1Wc1lYUnBiMjV6TVJNd0VRWURWUVFLREFwQmNIQnNaU0JKYm1NdU1Rc3dDUVlEVlFRR0V3SlZVekJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCT29UY2FQY3BlaXBOTDllUTA2dEN1N3BVY3dkQ1hkTjh2R3FhVWpkNThaOHRMeGlVQzBkQmVBK2V1TVlnZ2gxLzVpQWsrRk14VUZtQTJhMXI0YUNaOFNqZ2dJSU1JSUNCREFNQmdOVkhSTUJBZjhFQWpBQU1COEdBMVVkSXdRWU1CYUFGRDh2bENOUjAxREptaWc5N2JCODVjK2xrR0taTUhBR0NDc0dBUVVGQndFQkJHUXdZakF0QmdnckJnRUZCUWN3QW9ZaGFIUjBjRG92TDJObGNuUnpMbUZ3Y0d4bExtTnZiUzkzZDJSeVp6WXVaR1Z5TURFR0NDc0dBUVVGQnpBQmhpVm9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQXpMWGQzWkhKbk5qQXlNSUlCSGdZRFZSMGdCSUlCRlRDQ0FSRXdnZ0VOQmdvcWhraUc5Mk5rQlFZQk1JSCtNSUhEQmdnckJnRUZCUWNDQWpDQnRneUJzMUpsYkdsaGJtTmxJRzl1SUhSb2FYTWdZMlZ5ZEdsbWFXTmhkR1VnWW5rZ1lXNTVJSEJoY25SNUlHRnpjM1Z0WlhNZ1lXTmpaWEIwWVc1alpTQnZaaUIwYUdVZ2RHaGxiaUJoY0hCc2FXTmhZbXhsSUhOMFlXNWtZWEprSUhSbGNtMXpJR0Z1WkNCamIyNWthWFJwYjI1eklHOW1JSFZ6WlN3Z1kyVnlkR2xtYVdOaGRHVWdjRzlzYVdONUlHRnVaQ0JqWlhKMGFXWnBZMkYwYVc5dUlIQnlZV04wYVdObElITjBZWFJsYldWdWRITXVNRFlHQ0NzR0FRVUZCd0lCRmlwb2RIUndPaTh2ZDNkM0xtRndjR3hsTG1OdmJTOWpaWEowYVdacFkyRjBaV0YxZEdodmNtbDBlUzh3SFFZRFZSME9CQllFRkNPQ21NQnEvLzFMNWltdlZtcVgxb0NZZXFyTU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBUUJnb3Foa2lHOTJOa0Jnc0JCQUlGQURBS0JnZ3Foa2pPUFFRREF3Tm9BREJsQWpFQWw0SkI5R0pIaXhQMm51aWJ5VTFrM3dyaTVwc0dJeFBNRTA1c0ZLcTdoUXV6dmJleUJ1ODJGb3p6eG1ienBvZ29BakJMU0ZsMGRaV0lZbDJlalBWK0RpNWZCbktQdThteW1CUXRvRS9IMmJFUzBxQXM4Yk51ZVUzQ0JqamgxbHduRHNJPSIsIk1JSURGakNDQXB5Z0F3SUJBZ0lVSXNHaFJ3cDBjMm52VTRZU3ljYWZQVGp6Yk5jd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NakV3TXpFM01qQXpOekV3V2hjTk16WXdNekU1TURBd01EQXdXakIxTVVRd1FnWURWUVFERER0QmNIQnNaU0JYYjNKc1pIZHBaR1VnUkdWMlpXeHZjR1Z5SUZKbGJHRjBhVzl1Y3lCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVMTUFrR0ExVUVDd3dDUnpZeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUVic1FLQzk0UHJsV21aWG5YZ3R4emRWSkw4VDBTR1luZ0RSR3BuZ24zTjZQVDhKTUViN0ZEaTRiQm1QaENuWjMvc3E2UEYvY0djS1hXc0w1dk90ZVJoeUo0NXgzQVNQN2NPQithYW85MGZjcHhTdi9FWkZibmlBYk5nWkdoSWhwSW80SDZNSUgzTUJJR0ExVWRFd0VCL3dRSU1BWUJBZjhDQVFBd0h3WURWUjBqQkJnd0ZvQVV1N0Rlb1ZnemlKcWtpcG5ldnIzcnI5ckxKS3N3UmdZSUt3WUJCUVVIQVFFRU9qQTRNRFlHQ0NzR0FRVUZCekFCaGlwb2RIUndPaTh2YjJOemNDNWhjSEJzWlM1amIyMHZiMk56Y0RBekxXRndjR3hsY205dmRHTmhaek13TndZRFZSMGZCREF3TGpBc29DcWdLSVltYUhSMGNEb3ZMMk55YkM1aGNIQnNaUzVqYjIwdllYQndiR1Z5YjI5MFkyRm5NeTVqY213d0hRWURWUjBPQkJZRUZEOHZsQ05SMDFESm1pZzk3YkI4NWMrbGtHS1pNQTRHQTFVZER3RUIvd1FFQXdJQkJqQVFCZ29xaGtpRzkyTmtCZ0lCQkFJRkFEQUtCZ2dxaGtqT1BRUURBd05vQURCbEFqQkFYaFNxNUl5S29nTUNQdHc0OTBCYUI2NzdDYUVHSlh1ZlFCL0VxWkdkNkNTamlDdE9udU1UYlhWWG14eGN4ZmtDTVFEVFNQeGFyWlh2TnJreFUzVGtVTUkzM3l6dkZWVlJUNHd4V0pDOTk0T3NkY1o0K1JHTnNZRHlSNWdtZHIwbkRHZz0iLCJNSUlDUXpDQ0FjbWdBd0lCQWdJSUxjWDhpTkxGUzVVd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NVFF3TkRNd01UZ3hPVEEyV2hjTk16a3dORE13TVRneE9UQTJXakJuTVJzd0dRWURWUVFEREJKQmNIQnNaU0JTYjI5MElFTkJJQzBnUnpNeEpqQWtCZ05WQkFzTUhVRndjR3hsSUVObGNuUnBabWxqWVhScGIyNGdRWFYwYUc5eWFYUjVNUk13RVFZRFZRUUtEQXBCY0hCc1pTQkpibU11TVFzd0NRWURWUVFHRXdKVlV6QjJNQkFHQnlxR1NNNDlBZ0VHQlN1QkJBQWlBMklBQkpqcEx6MUFjcVR0a3lKeWdSTWMzUkNWOGNXalRuSGNGQmJaRHVXbUJTcDNaSHRmVGpqVHV4eEV0WC8xSDdZeVlsM0o2WVJiVHpCUEVWb0EvVmhZREtYMUR5eE5CMGNUZGRxWGw1ZHZNVnp0SzUxN0lEdll1VlRaWHBta09sRUtNYU5DTUVBd0hRWURWUjBPQkJZRUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdEZ1lEVlIwUEFRSC9CQVFEQWdFR01Bb0dDQ3FHU000OUJBTURBMmdBTUdVQ01RQ0Q2Y0hFRmw0YVhUUVkyZTN2OUd3T0FFWkx1Tit5UmhIRkQvM21lb3locG12T3dnUFVuUFdUeG5TNGF0K3FJeFVDTUcxbWloREsxQTNVVDgyTlF6NjBpbU9sTTI3amJkb1h0MlFmeUZNbStZaGlkRGtMRjF2TFVhZ002QmdENTZLeUtBPT0iXX0.eyJub3RpZmljYXRpb25UeXBlIjoiRElEX1JFTkVXIiwibm90aWZpY2F0aW9uVVVJRCI6IjQ2OWJmMzBlLTc3MTUtNGY5Zi1hYWUzLWE3YmZjMTJhZWE3NyIsImRhdGEiOnsiYnVuZGxlSWQiOiJjb20uYXVkYW9zLmF1ZGFyZWNvcmRlciIsImJ1bmRsZVZlcnNpb24iOiI3IiwiZW52aXJvbm1lbnQiOiJTYW5kYm94Iiwic2lnbmVkVHJhbnNhY3Rpb25JbmZvIjoiZXlKaGJHY2lPaUpGVXpJMU5pSXNJbmcxWXlJNld5Sk5TVWxGVFVSRFEwRTNZV2RCZDBsQ1FXZEpVV0ZRYjFCc1pIWndVMjlGU0RCc1FuSnFSRkIyT1dwQlMwSm5aM0ZvYTJwUFVGRlJSRUY2UWpGTlZWRjNVV2RaUkZaUlVVUkVSSFJDWTBoQ2MxcFRRbGhpTTBweldraGtjRnBIVldkU1IxWXlXbGQ0ZG1OSFZubEpSa3BzWWtkR01HRlhPWFZqZVVKRVdsaEtNR0ZYV25CWk1rWXdZVmM1ZFVsRlJqRmtSMmgyWTIxc01HVlVSVXhOUVd0SFFURlZSVU4zZDBOU2VsbDRSWHBCVWtKblRsWkNRVzlOUTJ0R2QyTkhlR3hKUld4MVdYazBlRU42UVVwQ1owNVdRa0ZaVkVGc1ZsUk5RalJZUkZSSmVFMUVaM2xPVkVGNVRsUkJlazVHYjFoRVZFbDZUVVJyZVU1RVFYbE9WRUY2VFRGdmQyZGFTWGhSUkVFclFtZE9Wa0pCVFUxT01VSjVZakpSWjFKVlRrUkpSVEZvV1hsQ1FtTklRV2RWTTFKMlkyMVZaMWxYTld0SlIyeFZaRmMxYkdONVFsUmtSemw1V2xOQ1UxcFhUbXhoV0VJd1NVWk9jRm95TlhCaWJXTjRURVJCY1VKblRsWkNRWE5OU1RCR2QyTkhlR3hKUm1SMlkyMTRhMlF5Ykd0YVUwSkZXbGhhYkdKSE9YZGFXRWxuVlcxV2MxbFlVbkJpTWpWNlRWSk5kMFZSV1VSV1VWRkxSRUZ3UW1OSVFuTmFVMEpLWW0xTmRVMVJjM2REVVZsRVZsRlJSMFYzU2xaVmVrSmFUVUpOUjBKNWNVZFRUVFE1UVdkRlIwTkRjVWRUVFRRNVFYZEZTRUV3U1VGQ1QyOVVZMkZRWTNCbGFYQk9URGxsVVRBMmRFTjFOM0JWWTNka1ExaGtUamgyUjNGaFZXcGtOVGhhT0hSTWVHbFZRekJrUW1WQksyVjFUVmxuWjJneEx6VnBRV3NyUmsxNFZVWnRRVEpoTVhJMFlVTmFPRk5xWjJkSlNVMUpTVU5DUkVGTlFtZE9Wa2hTVFVKQlpqaEZRV3BCUVUxQ09FZEJNVlZrU1hkUldVMUNZVUZHUkRoMmJFTk9VakF4UkVwdGFXYzVOMkpDT0RWaksyeHJSMHRhVFVoQlIwTkRjMGRCVVZWR1FuZEZRa0pIVVhkWmFrRjBRbWRuY2tKblJVWkNVV04zUVc5WmFHRklVakJqUkc5MlRESk9iR051VW5wTWJVWjNZMGQ0YkV4dFRuWmlVemt6WkRKU2VWcDZXWFZhUjFaNVRVUkZSME5EYzBkQlVWVkdRbnBCUW1ocFZtOWtTRkozVDJrNGRtSXlUbnBqUXpWb1kwaENjMXBUTldwaU1qQjJZakpPZW1ORVFYcE1XR1F6V2toS2JrNXFRWGxOU1VsQ1NHZFpSRlpTTUdkQ1NVbENSbFJEUTBGU1JYZG5aMFZPUW1kdmNXaHJhVWM1TWs1clFsRlpRazFKU0N0TlNVaEVRbWRuY2tKblJVWkNVV05EUVdwRFFuUm5lVUp6TVVwc1lrZHNhR0p0VG14SlJ6bDFTVWhTYjJGWVRXZFpNbFo1WkVkc2JXRlhUbWhrUjFWbldXNXJaMWxYTlRWSlNFSm9ZMjVTTlVsSFJucGpNMVowV2xoTloxbFhUbXBhV0VJd1dWYzFhbHBUUW5aYWFVSXdZVWRWWjJSSGFHeGlhVUpvWTBoQ2MyRlhUbWhaYlhoc1NVaE9NRmxYTld0WldFcHJTVWhTYkdOdE1YcEpSMFoxV2tOQ2FtSXlOV3RoV0ZKd1lqSTFla2xIT1cxSlNGWjZXbE4zWjFreVZubGtSMnh0WVZkT2FHUkhWV2RqUnpsellWZE9OVWxIUm5WYVEwSnFXbGhLTUdGWFduQlpNa1l3WVZjNWRVbElRbmxaVjA0d1lWZE9iRWxJVGpCWldGSnNZbGRXZFdSSVRYVk5SRmxIUTBOelIwRlJWVVpDZDBsQ1JtbHdiMlJJVW5kUGFUaDJaRE5rTTB4dFJuZGpSM2hzVEcxT2RtSlRPV3BhV0Vvd1lWZGFjRmt5UmpCYVYwWXhaRWRvZG1OdGJEQmxVemgzU0ZGWlJGWlNNRTlDUWxsRlJrTlBRMjFOUW5Fdkx6Rk1OV2x0ZGxadGNWZ3hiME5aWlhGeVRVMUJORWRCTVZWa1JIZEZRaTkzVVVWQmQwbElaMFJCVVVKbmIzRm9hMmxIT1RKT2EwSm5jMEpDUVVsR1FVUkJTMEpuWjNGb2EycFBVRkZSUkVGM1RtOUJSRUpzUVdwRlFXdzBTa0k1UjBwSWFYaFFNbTUxYVdKNVZURnJNM2R5YVRWd2MwZEplRkJOUlRBMWMwWkxjVGRvVVhWNmRtSmxlVUoxT0RKR2IzcDZlRzFpZW5CdloyOUJha0pNVTBac01HUmFWMGxaYkRKbGFsQldLMFJwTldaQ2JrdFFkVGh0ZVcxQ1VYUnZSUzlJTW1KRlV6QnhRWE00WWs1MVpWVXpRMEpxYW1neGJIZHVSSE5KUFNJc0lrMUpTVVJHYWtORFFYQjVaMEYzU1VKQlowbFZTWE5IYUZKM2NEQmpNbTUyVlRSWlUzbGpZV1pRVkdwNllrNWpkME5uV1VsTGIxcEplbW93UlVGM1RYZGFla1ZpVFVKclIwRXhWVVZCZDNkVFVWaENkMkpIVldkVmJUbDJaRU5DUkZGVFFYUkpSV042VFZOWmQwcEJXVVJXVVZGTVJFSXhRbU5JUW5OYVUwSkVXbGhLTUdGWFduQlpNa1l3WVZjNWRVbEZSakZrUjJoMlkyMXNNR1ZVUlZSTlFrVkhRVEZWUlVObmQwdFJXRUozWWtkVloxTlhOV3BNYWtWTVRVRnJSMEV4VlVWQ2FFMURWbFpOZDBob1kwNU5ha1YzVFhwRk0wMXFRWHBPZWtWM1YyaGpUazE2V1hkTmVrVTFUVVJCZDAxRVFYZFhha0l4VFZWUmQxRm5XVVJXVVZGRVJFUjBRbU5JUW5OYVUwSllZak5LYzFwSVpIQmFSMVZuVWtkV01scFhlSFpqUjFaNVNVWktiR0pIUmpCaFZ6bDFZM2xDUkZwWVNqQmhWMXB3V1RKR01HRlhPWFZKUlVZeFpFZG9kbU50YkRCbFZFVk1UVUZyUjBFeFZVVkRkM2REVW5wWmVFVjZRVkpDWjA1V1FrRnZUVU5yUm5kalIzaHNTVVZzZFZsNU5IaERla0ZLUW1kT1ZrSkJXVlJCYkZaVVRVaFpkMFZCV1VoTGIxcEplbW93UTBGUldVWkxORVZGUVVOSlJGbG5RVVZpYzFGTFF6azBVSEpzVjIxYVdHNVlaM1I0ZW1SV1NrdzRWREJUUjFsdVowUlNSM0J1WjI0elRqWlFWRGhLVFVWaU4wWkVhVFJpUW0xUWFFTnVXak12YzNFMlVFWXZZMGRqUzFoWGMwdzFkazkwWlZKb2VVbzBOWGd6UVZOUU4yTlBRaXRoWVc4NU1HWmpjSGhUZGk5RldrWmlibWxCWWs1bldrZG9TV2h3U1c4MFNEWk5TVWd6VFVKSlIwRXhWV1JGZDBWQ0wzZFJTVTFCV1VKQlpqaERRVkZCZDBoM1dVUldVakJxUWtKbmQwWnZRVlYxTjBSbGIxWm5lbWxLY1d0cGNHNWxkbkl6Y25JNWNreEtTM04zVW1kWlNVdDNXVUpDVVZWSVFWRkZSVTlxUVRSTlJGbEhRME56UjBGUlZVWkNla0ZDYUdsd2IyUklVbmRQYVRoMllqSk9lbU5ETldoalNFSnpXbE0xYW1JeU1IWmlNazU2WTBSQmVreFhSbmRqUjNoc1kyMDVkbVJIVG1oYWVrMTNUbmRaUkZaU01HWkNSRUYzVEdwQmMyOURjV2RMU1ZsdFlVaFNNR05FYjNaTU1rNTVZa00xYUdOSVFuTmFVelZxWWpJd2RsbFlRbmRpUjFaNVlqSTVNRmt5Um01TmVUVnFZMjEzZDBoUldVUldVakJQUWtKWlJVWkVPSFpzUTA1U01ERkVTbTFwWnprM1lrSTROV01yYkd0SFMxcE5RVFJIUVRGVlpFUjNSVUl2ZDFGRlFYZEpRa0pxUVZGQ1oyOXhhR3RwUnpreVRtdENaMGxDUWtGSlJrRkVRVXRDWjJkeGFHdHFUMUJSVVVSQmQwNXZRVVJDYkVGcVFrRllhRk54TlVsNVMyOW5UVU5RZEhjME9UQkNZVUkyTnpkRFlVVkhTbGgxWmxGQ0wwVnhXa2RrTmtOVGFtbERkRTl1ZFUxVVlsaFdXRzE0ZUdONFptdERUVkZFVkZOUWVHRnlXbGgyVG5KcmVGVXpWR3RWVFVrek0zbDZka1pXVmxKVU5IZDRWMHBET1RrMFQzTmtZMW8wSzFKSFRuTlpSSGxTTldkdFpISXdia1JIWnowaUxDSk5TVWxEVVhwRFEwRmpiV2RCZDBsQ1FXZEpTVXhqV0RocFRreEdVelZWZDBObldVbExiMXBKZW1vd1JVRjNUWGRhZWtWaVRVSnJSMEV4VlVWQmQzZFRVVmhDZDJKSFZXZFZiVGwyWkVOQ1JGRlRRWFJKUldONlRWTlpkMHBCV1VSV1VWRk1SRUl4UW1OSVFuTmFVMEpFV2xoS01HRlhXbkJaTWtZd1lWYzVkVWxGUmpGa1IyaDJZMjFzTUdWVVJWUk5Ra1ZIUVRGVlJVTm5kMHRSV0VKM1lrZFZaMU5YTldwTWFrVk1UVUZyUjBFeFZVVkNhRTFEVmxaTmQwaG9ZMDVOVkZGM1RrUk5kMDFVWjNoUFZFRXlWMmhqVGsxNmEzZE9SRTEzVFZSbmVFOVVRVEpYYWtKdVRWSnpkMGRSV1VSV1VWRkVSRUpLUW1OSVFuTmFVMEpUWWpJNU1FbEZUa0pKUXpCblVucE5lRXBxUVd0Q1owNVdRa0Z6VFVoVlJuZGpSM2hzU1VWT2JHTnVVbkJhYld4cVdWaFNjR0l5TkdkUldGWXdZVWM1ZVdGWVVqVk5VazEzUlZGWlJGWlJVVXRFUVhCQ1kwaENjMXBUUWtwaWJVMTFUVkZ6ZDBOUldVUldVVkZIUlhkS1ZsVjZRakpOUWtGSFFubHhSMU5OTkRsQlowVkhRbE4xUWtKQlFXbEJNa2xCUWtwcWNFeDZNVUZqY1ZSMGEzbEtlV2RTVFdNelVrTldPR05YYWxSdVNHTkdRbUphUkhWWGJVSlRjRE5hU0hSbVZHcHFWSFY0ZUVWMFdDOHhTRGRaZVZsc00wbzJXVkppVkhwQ1VFVldiMEV2Vm1oWlJFdFlNVVI1ZUU1Q01HTlVaR1J4V0d3MVpIWk5WbnAwU3pVeE4wbEVkbGwxVmxSYVdIQnRhMDlzUlV0TllVNURUVVZCZDBoUldVUldVakJQUWtKWlJVWk1kWGN6Y1VaWlRUUnBZWEJKY1ZvemNqWTVOall2WVhsNVUzSk5RVGhIUVRGVlpFVjNSVUl2ZDFGR1RVRk5Ra0ZtT0hkRVoxbEVWbEl3VUVGUlNDOUNRVkZFUVdkRlIwMUJiMGREUTNGSFUwMDBPVUpCVFVSQk1tZEJUVWRWUTAxUlEwUTJZMGhGUm13MFlWaFVVVmt5WlROMk9VZDNUMEZGV2t4MVRpdDVVbWhJUmtRdk0yMWxiM2xvY0cxMlQzZG5VRlZ1VUZkVWVHNVROR0YwSzNGSmVGVkRUVWN4Yldsb1JFc3hRVE5WVkRneVRsRjZOakJwYlU5c1RUSTNhbUprYjFoME1sRm1lVVpOYlN0WmFHbGtSR3RNUmpGMlRGVmhaMDAyUW1kRU5UWkxlVXRCUFQwaVhYMC5leUowY21GdWMyRmpkR2x2Ymtsa0lqb2lNakF3TURBd01EQXdOREEwTnpFeE9TSXNJbTl5YVdkcGJtRnNWSEpoYm5OaFkzUnBiMjVKWkNJNklqSXdNREF3TURBd01EQTROREkyTURjaUxDSjNaV0pQY21SbGNreHBibVZKZEdWdFNXUWlPaUl5TURBd01EQXdNREF3TXpBeU9ETXlJaXdpWW5WdVpHeGxTV1FpT2lKamIyMHVZWFZrWVc5ekxtRjFaR0Z5WldOdmNtUmxjaUlzSW5CeWIyUjFZM1JKWkNJNkltTnZiUzVoZFdSaGIzTXVZWFZrWVhKbFkyOXlaR1Z5TG5acGNDNXRNaUlzSW5OMVluTmpjbWx3ZEdsdmJrZHliM1Z3U1dSbGJuUnBabWxsY2lJNklqSXdPVEk1TlRNMklpd2ljSFZ5WTJoaGMyVkVZWFJsSWpveE5qUTJNemczTURFMk1EQXdMQ0p2Y21sbmFXNWhiRkIxY21Ob1lYTmxSR0YwWlNJNk1UWTBOakEwTmpBek56QXdNQ3dpWlhod2FYSmxjMFJoZEdVaU9qRTJORFl6T0RjeE9UWXdNREFzSW5GMVlXNTBhWFI1SWpveExDSjBlWEJsSWpvaVFYVjBieTFTWlc1bGQyRmliR1VnVTNWaWMyTnlhWEIwYVc5dUlpd2lhVzVCY0hCUGQyNWxjbk5vYVhCVWVYQmxJam9pVUZWU1EwaEJVMFZFSWl3aWMybG5ibVZrUkdGMFpTSTZNVFkwTmpNNE56QXdPREkxTkN3aVpXNTJhWEp2Ym0xbGJuUWlPaUpUWVc1a1ltOTRJbjAubXZWbEFDRENjUy1sMW9BWVJJOUpWc3kzZE9RMHNuamRBN1FsYWcxNURyUThsbUFvc3JranJCUlB6T25UQ2FOVWtVZ3VzT01kQkREUGZoZm9TazBvTXciLCJzaWduZWRSZW5ld2FsSW5mbyI6ImV5SmhiR2NpT2lKRlV6STFOaUlzSW5nMVl5STZXeUpOU1VsRlRVUkRRMEUzWVdkQmQwbENRV2RKVVdGUWIxQnNaSFp3VTI5RlNEQnNRbkpxUkZCMk9XcEJTMEpuWjNGb2EycFBVRkZSUkVGNlFqRk5WVkYzVVdkWlJGWlJVVVJFUkhSQ1kwaENjMXBUUWxoaU0wcHpXa2hrY0ZwSFZXZFNSMVl5V2xkNGRtTkhWbmxKUmtwc1lrZEdNR0ZYT1hWamVVSkVXbGhLTUdGWFduQlpNa1l3WVZjNWRVbEZSakZrUjJoMlkyMXNNR1ZVUlV4TlFXdEhRVEZWUlVOM2QwTlNlbGw0UlhwQlVrSm5UbFpDUVc5TlEydEdkMk5IZUd4SlJXeDFXWGswZUVONlFVcENaMDVXUWtGWlZFRnNWbFJOUWpSWVJGUkplRTFFWjNsT1ZFRjVUbFJCZWs1R2IxaEVWRWw2VFVScmVVNUVRWGxPVkVGNlRURnZkMmRhU1hoUlJFRXJRbWRPVmtKQlRVMU9NVUo1WWpKUloxSlZUa1JKUlRGb1dYbENRbU5JUVdkVk0xSjJZMjFWWjFsWE5XdEpSMnhWWkZjMWJHTjVRbFJrUnpsNVdsTkNVMXBYVG14aFdFSXdTVVpPY0ZveU5YQmliV040VEVSQmNVSm5UbFpDUVhOTlNUQkdkMk5IZUd4SlJtUjJZMjE0YTJReWJHdGFVMEpGV2xoYWJHSkhPWGRhV0VsblZXMVdjMWxZVW5CaU1qVjZUVkpOZDBWUldVUldVVkZMUkVGd1FtTklRbk5hVTBKS1ltMU5kVTFSYzNkRFVWbEVWbEZSUjBWM1NsWlZla0phVFVKTlIwSjVjVWRUVFRRNVFXZEZSME5EY1VkVFRUUTVRWGRGU0VFd1NVRkNUMjlVWTJGUVkzQmxhWEJPVERsbFVUQTJkRU4xTjNCVlkzZGtRMWhrVGpoMlIzRmhWV3BrTlRoYU9IUk1lR2xWUXpCa1FtVkJLMlYxVFZsbloyZ3hMelZwUVdzclJrMTRWVVp0UVRKaE1YSTBZVU5hT0ZOcVoyZEpTVTFKU1VOQ1JFRk5RbWRPVmtoU1RVSkJaamhGUVdwQlFVMUNPRWRCTVZWa1NYZFJXVTFDWVVGR1JEaDJiRU5PVWpBeFJFcHRhV2M1TjJKQ09EVmpLMnhyUjB0YVRVaEJSME5EYzBkQlVWVkdRbmRGUWtKSFVYZFpha0YwUW1kbmNrSm5SVVpDVVdOM1FXOVphR0ZJVWpCalJHOTJUREpPYkdOdVVucE1iVVozWTBkNGJFeHRUblppVXprelpESlNlVnA2V1hWYVIxWjVUVVJGUjBORGMwZEJVVlZHUW5wQlFtaHBWbTlrU0ZKM1QyazRkbUl5VG5walF6Vm9ZMGhDYzFwVE5XcGlNakIyWWpKT2VtTkVRWHBNV0dReldraEtiazVxUVhsTlNVbENTR2RaUkZaU01HZENTVWxDUmxSRFEwRlNSWGRuWjBWT1FtZHZjV2hyYVVjNU1rNXJRbEZaUWsxSlNDdE5TVWhFUW1kbmNrSm5SVVpDVVdORFFXcERRblJuZVVKek1VcHNZa2RzYUdKdFRteEpSemwxU1VoU2IyRllUV2RaTWxaNVpFZHNiV0ZYVG1oa1IxVm5XVzVyWjFsWE5UVkpTRUpvWTI1U05VbEhSbnBqTTFaMFdsaE5aMWxYVG1wYVdFSXdXVmMxYWxwVFFuWmFhVUl3WVVkVloyUkhhR3hpYVVKb1kwaENjMkZYVG1oWmJYaHNTVWhPTUZsWE5XdFpXRXByU1VoU2JHTnRNWHBKUjBaMVdrTkNhbUl5Tld0aFdGSndZakkxZWtsSE9XMUpTRlo2V2xOM1oxa3lWbmxrUjJ4dFlWZE9hR1JIVldkalJ6bHpZVmRPTlVsSFJuVmFRMEpxV2xoS01HRlhXbkJaTWtZd1lWYzVkVWxJUW5sWlYwNHdZVmRPYkVsSVRqQlpXRkpzWWxkV2RXUklUWFZOUkZsSFEwTnpSMEZSVlVaQ2QwbENSbWx3YjJSSVVuZFBhVGgyWkROa00weHRSbmRqUjNoc1RHMU9kbUpUT1dwYVdFb3dZVmRhY0ZreVJqQmFWMFl4WkVkb2RtTnRiREJsVXpoM1NGRlpSRlpTTUU5Q1FsbEZSa05QUTIxTlFuRXZMekZNTldsdGRsWnRjVmd4YjBOWlpYRnlUVTFCTkVkQk1WVmtSSGRGUWk5M1VVVkJkMGxJWjBSQlVVSm5iM0ZvYTJsSE9USk9hMEpuYzBKQ1FVbEdRVVJCUzBKblozRm9hMnBQVUZGUlJFRjNUbTlCUkVKc1FXcEZRV3cwU2tJNVIwcElhWGhRTW01MWFXSjVWVEZyTTNkeWFUVndjMGRKZUZCTlJUQTFjMFpMY1Rkb1VYVjZkbUpsZVVKMU9ESkdiM3A2ZUcxaWVuQnZaMjlCYWtKTVUwWnNNR1JhVjBsWmJESmxhbEJXSzBScE5XWkNia3RRZFRodGVXMUNVWFJ2UlM5SU1tSkZVekJ4UVhNNFlrNTFaVlV6UTBKcWFtZ3hiSGR1UkhOSlBTSXNJazFKU1VSR2FrTkRRWEI1WjBGM1NVSkJaMGxWU1hOSGFGSjNjREJqTW01MlZUUlpVM2xqWVdaUVZHcDZZazVqZDBObldVbExiMXBKZW1vd1JVRjNUWGRhZWtWaVRVSnJSMEV4VlVWQmQzZFRVVmhDZDJKSFZXZFZiVGwyWkVOQ1JGRlRRWFJKUldONlRWTlpkMHBCV1VSV1VWRk1SRUl4UW1OSVFuTmFVMEpFV2xoS01HRlhXbkJaTWtZd1lWYzVkVWxGUmpGa1IyaDJZMjFzTUdWVVJWUk5Ra1ZIUVRGVlJVTm5kMHRSV0VKM1lrZFZaMU5YTldwTWFrVk1UVUZyUjBFeFZVVkNhRTFEVmxaTmQwaG9ZMDVOYWtWM1RYcEZNMDFxUVhwT2VrVjNWMmhqVGsxNldYZE5la1UxVFVSQmQwMUVRWGRYYWtJeFRWVlJkMUZuV1VSV1VWRkVSRVIwUW1OSVFuTmFVMEpZWWpOS2MxcElaSEJhUjFWblVrZFdNbHBYZUhaalIxWjVTVVpLYkdKSFJqQmhWemwxWTNsQ1JGcFlTakJoVjFwd1dUSkdNR0ZYT1hWSlJVWXhaRWRvZG1OdGJEQmxWRVZNVFVGclIwRXhWVVZEZDNkRFVucFplRVY2UVZKQ1owNVdRa0Z2VFVOclJuZGpSM2hzU1VWc2RWbDVOSGhEZWtGS1FtZE9Wa0pCV1ZSQmJGWlVUVWhaZDBWQldVaExiMXBKZW1vd1EwRlJXVVpMTkVWRlFVTkpSRmxuUVVWaWMxRkxRemswVUhKc1YyMWFXRzVZWjNSNGVtUldTa3c0VkRCVFIxbHVaMFJTUjNCdVoyNHpUalpRVkRoS1RVVmlOMFpFYVRSaVFtMVFhRU51V2pNdmMzRTJVRVl2WTBkalMxaFhjMHcxZGs5MFpWSm9lVW8wTlhnelFWTlFOMk5QUWl0aFlXODVNR1pqY0hoVGRpOUZXa1ppYm1sQllrNW5Xa2RvU1dod1NXODBTRFpOU1VnelRVSkpSMEV4VldSRmQwVkNMM2RSU1UxQldVSkJaamhEUVZGQmQwaDNXVVJXVWpCcVFrSm5kMFp2UVZWMU4wUmxiMVpuZW1sS2NXdHBjRzVsZG5JemNuSTVja3hLUzNOM1VtZFpTVXQzV1VKQ1VWVklRVkZGUlU5cVFUUk5SRmxIUTBOelIwRlJWVVpDZWtGQ2FHbHdiMlJJVW5kUGFUaDJZakpPZW1ORE5XaGpTRUp6V2xNMWFtSXlNSFppTWs1NlkwUkJla3hYUm5kalIzaHNZMjA1ZG1SSFRtaGFlazEzVG5kWlJGWlNNR1pDUkVGM1RHcEJjMjlEY1dkTFNWbHRZVWhTTUdORWIzWk1NazU1WWtNMWFHTklRbk5hVXpWcVlqSXdkbGxZUW5kaVIxWjVZakk1TUZreVJtNU5lVFZxWTIxM2QwaFJXVVJXVWpCUFFrSlpSVVpFT0hac1EwNVNNREZFU20xcFp6azNZa0k0TldNcmJHdEhTMXBOUVRSSFFURlZaRVIzUlVJdmQxRkZRWGRKUWtKcVFWRkNaMjl4YUd0cFJ6a3lUbXRDWjBsQ1FrRkpSa0ZFUVV0Q1oyZHhhR3RxVDFCUlVVUkJkMDV2UVVSQ2JFRnFRa0ZZYUZOeE5VbDVTMjluVFVOUWRIYzBPVEJDWVVJMk56ZERZVVZIU2xoMVpsRkNMMFZ4V2tka05rTlRhbWxEZEU5dWRVMVVZbGhXV0cxNGVHTjRabXREVFZGRVZGTlFlR0Z5V2xoMlRuSnJlRlV6Vkd0VlRVa3pNM2w2ZGtaV1ZsSlVOSGQ0VjBwRE9UazBUM05rWTFvMEsxSkhUbk5aUkhsU05XZHRaSEl3YmtSSFp6MGlMQ0pOU1VsRFVYcERRMEZqYldkQmQwbENRV2RKU1V4aldEaHBUa3hHVXpWVmQwTm5XVWxMYjFwSmVtb3dSVUYzVFhkYWVrVmlUVUpyUjBFeFZVVkJkM2RUVVZoQ2QySkhWV2RWYlRsMlpFTkNSRkZUUVhSSlJXTjZUVk5aZDBwQldVUldVVkZNUkVJeFFtTklRbk5hVTBKRVdsaEtNR0ZYV25CWk1rWXdZVmM1ZFVsRlJqRmtSMmgyWTIxc01HVlVSVlJOUWtWSFFURlZSVU5uZDB0UldFSjNZa2RWWjFOWE5XcE1ha1ZNVFVGclIwRXhWVVZDYUUxRFZsWk5kMGhvWTA1TlZGRjNUa1JOZDAxVVozaFBWRUV5VjJoalRrMTZhM2RPUkUxM1RWUm5lRTlVUVRKWGFrSnVUVkp6ZDBkUldVUldVVkZFUkVKS1FtTklRbk5hVTBKVFlqSTVNRWxGVGtKSlF6Qm5VbnBOZUVwcVFXdENaMDVXUWtGelRVaFZSbmRqUjNoc1NVVk9iR051VW5CYWJXeHFXVmhTY0dJeU5HZFJXRll3WVVjNWVXRllValZOVWsxM1JWRlpSRlpSVVV0RVFYQkNZMGhDYzFwVFFrcGliVTExVFZGemQwTlJXVVJXVVZGSFJYZEtWbFY2UWpKTlFrRkhRbmx4UjFOTk5EbEJaMFZIUWxOMVFrSkJRV2xCTWtsQlFrcHFjRXg2TVVGamNWUjBhM2xLZVdkU1RXTXpVa05XT0dOWGFsUnVTR05HUW1KYVJIVlhiVUpUY0ROYVNIUm1WR3BxVkhWNGVFVjBXQzh4U0RkWmVWbHNNMG8yV1ZKaVZIcENVRVZXYjBFdlZtaFpSRXRZTVVSNWVFNUNNR05VWkdSeFdHdzFaSFpOVm5wMFN6VXhOMGxFZGxsMVZsUmFXSEJ0YTA5c1JVdE5ZVTVEVFVWQmQwaFJXVVJXVWpCUFFrSlpSVVpNZFhjemNVWlpUVFJwWVhCSmNWb3pjalk1TmpZdllYbDVVM0pOUVRoSFFURlZaRVYzUlVJdmQxRkdUVUZOUWtGbU9IZEVaMWxFVmxJd1VFRlJTQzlDUVZGRVFXZEZSMDFCYjBkRFEzRkhVMDAwT1VKQlRVUkJNbWRCVFVkVlEwMVJRMFEyWTBoRlJtdzBZVmhVVVZreVpUTjJPVWQzVDBGRldreDFUaXQ1VW1oSVJrUXZNMjFsYjNsb2NHMTJUM2RuVUZWdVVGZFVlRzVUTkdGMEszRkplRlZEVFVjeGJXbG9SRXN4UVROVlZEZ3lUbEY2TmpCcGJVOXNUVEkzYW1Ka2IxaDBNbEZtZVVaTmJTdFphR2xrUkd0TVJqRjJURlZoWjAwMlFtZEVOVFpMZVV0QlBUMGlYWDAuZXlKdmNtbG5hVzVoYkZSeVlXNXpZV04wYVc5dVNXUWlPaUl5TURBd01EQXdNREF3T0RReU5qQTNJaXdpWVhWMGIxSmxibVYzVUhKdlpIVmpkRWxrSWpvaVkyOXRMbUYxWkdGdmN5NWhkV1JoY21WamIzSmtaWEl1ZG1sd0xtMHlJaXdpY0hKdlpIVmpkRWxrSWpvaVkyOXRMbUYxWkdGdmN5NWhkV1JoY21WamIzSmtaWEl1ZG1sd0xtMHlJaXdpWVhWMGIxSmxibVYzVTNSaGRIVnpJam94TENKemFXZHVaV1JFWVhSbElqb3hOalEyTXpnM01EQTRNakk0TENKbGJuWnBjbTl1YldWdWRDSTZJbE5oYm1SaWIzZ2lmUS5qT3pWbkEzZU9kOHh1aThfcEppY3U5V1hWSFp6WksxMDRMTVRtcl9DNmVMdnRkaHpGVVEtWkRnY0IyRVRfY2xCZEc5S2tweFc3OEV3b21sai1hRmZPUSJ9LCJ2ZXJzaW9uIjoiMi4wIn0.AJ7k5AaTtnbX-n9rKwt1YuqtbUeFe8ah7C6TsCi4b1nfkl9ZdZVN4n8u63lDudzol4AzJw-cYpFvDew28KqyvA"}`

func TestDecodeSignedPayload(t *testing.T) {
	req := &NotificationV2Req{}
	err := json.Unmarshal([]byte(notificationV2Body), req)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Color(xlog.CyanBright).Infof("req: %+v", *req)
	// decode signedPayload
	payload, err := DecodeSignedPayload(req.SignedPayload)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("payload.NotificationType: %s", payload.NotificationType)
	xlog.Debugf("payload.Subtype: %s", payload.Subtype)
	xlog.Debugf("payload.NotificationUUID: %s", payload.NotificationUUID)
	xlog.Debugf("payload.NotificationVersion: %s", payload.NotificationVersion)
	xlog.Debugf("payload.Data: %+v", payload.Data)
	bs1, _ := json.Marshal(payload)
	xlog.Color(xlog.RedBright).Info(string(bs1))
	/*
		{
		    "notificationType":"DID_RENEW",
		    "subtype":"",
		    "notificationUUID":"469bf30e-7715-4f9f-aae3-a7bfc12aea77",
		    "notificationVersion":"",
		    "data":{
		        "appAppleId":0,
		        "bundleId":"com.audaos.audarecorder",
		        "bundleVersion":"7",
		        "environment":"Sandbox",
		        "signedRenewalInfo":"eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlFTURDQ0E3YWdBd0lCQWdJUWFQb1BsZHZwU29FSDBsQnJqRFB2OWpBS0JnZ3Foa2pPUFFRREF6QjFNVVF3UWdZRFZRUURERHRCY0hCc1pTQlhiM0pzWkhkcFpHVWdSR1YyWld4dmNHVnlJRkpsYkdGMGFXOXVjeUJEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURUxNQWtHQTFVRUN3d0NSell4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRJeE1EZ3lOVEF5TlRBek5Gb1hEVEl6TURreU5EQXlOVEF6TTFvd2daSXhRREErQmdOVkJBTU1OMUJ5YjJRZ1JVTkRJRTFoWXlCQmNIQWdVM1J2Y21VZ1lXNWtJR2xVZFc1bGN5QlRkRzl5WlNCU1pXTmxhWEIwSUZOcFoyNXBibWN4TERBcUJnTlZCQXNNSTBGd2NHeGxJRmR2Y214a2QybGtaU0JFWlhabGJHOXdaWElnVW1Wc1lYUnBiMjV6TVJNd0VRWURWUVFLREFwQmNIQnNaU0JKYm1NdU1Rc3dDUVlEVlFRR0V3SlZVekJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCT29UY2FQY3BlaXBOTDllUTA2dEN1N3BVY3dkQ1hkTjh2R3FhVWpkNThaOHRMeGlVQzBkQmVBK2V1TVlnZ2gxLzVpQWsrRk14VUZtQTJhMXI0YUNaOFNqZ2dJSU1JSUNCREFNQmdOVkhSTUJBZjhFQWpBQU1COEdBMVVkSXdRWU1CYUFGRDh2bENOUjAxREptaWc5N2JCODVjK2xrR0taTUhBR0NDc0dBUVVGQndFQkJHUXdZakF0QmdnckJnRUZCUWN3QW9ZaGFIUjBjRG92TDJObGNuUnpMbUZ3Y0d4bExtTnZiUzkzZDJSeVp6WXVaR1Z5TURFR0NDc0dBUVVGQnpBQmhpVm9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQXpMWGQzWkhKbk5qQXlNSUlCSGdZRFZSMGdCSUlCRlRDQ0FSRXdnZ0VOQmdvcWhraUc5Mk5rQlFZQk1JSCtNSUhEQmdnckJnRUZCUWNDQWpDQnRneUJzMUpsYkdsaGJtTmxJRzl1SUhSb2FYTWdZMlZ5ZEdsbWFXTmhkR1VnWW5rZ1lXNTVJSEJoY25SNUlHRnpjM1Z0WlhNZ1lXTmpaWEIwWVc1alpTQnZaaUIwYUdVZ2RHaGxiaUJoY0hCc2FXTmhZbXhsSUhOMFlXNWtZWEprSUhSbGNtMXpJR0Z1WkNCamIyNWthWFJwYjI1eklHOW1JSFZ6WlN3Z1kyVnlkR2xtYVdOaGRHVWdjRzlzYVdONUlHRnVaQ0JqWlhKMGFXWnBZMkYwYVc5dUlIQnlZV04wYVdObElITjBZWFJsYldWdWRITXVNRFlHQ0NzR0FRVUZCd0lCRmlwb2RIUndPaTh2ZDNkM0xtRndjR3hsTG1OdmJTOWpaWEowYVdacFkyRjBaV0YxZEdodmNtbDBlUzh3SFFZRFZSME9CQllFRkNPQ21NQnEvLzFMNWltdlZtcVgxb0NZZXFyTU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBUUJnb3Foa2lHOTJOa0Jnc0JCQUlGQURBS0JnZ3Foa2pPUFFRREF3Tm9BREJsQWpFQWw0SkI5R0pIaXhQMm51aWJ5VTFrM3dyaTVwc0dJeFBNRTA1c0ZLcTdoUXV6dmJleUJ1ODJGb3p6eG1ienBvZ29BakJMU0ZsMGRaV0lZbDJlalBWK0RpNWZCbktQdThteW1CUXRvRS9IMmJFUzBxQXM4Yk51ZVUzQ0JqamgxbHduRHNJPSIsIk1JSURGakNDQXB5Z0F3SUJBZ0lVSXNHaFJ3cDBjMm52VTRZU3ljYWZQVGp6Yk5jd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NakV3TXpFM01qQXpOekV3V2hjTk16WXdNekU1TURBd01EQXdXakIxTVVRd1FnWURWUVFERER0QmNIQnNaU0JYYjNKc1pIZHBaR1VnUkdWMlpXeHZjR1Z5SUZKbGJHRjBhVzl1Y3lCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVMTUFrR0ExVUVDd3dDUnpZeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUVic1FLQzk0UHJsV21aWG5YZ3R4emRWSkw4VDBTR1luZ0RSR3BuZ24zTjZQVDhKTUViN0ZEaTRiQm1QaENuWjMvc3E2UEYvY0djS1hXc0w1dk90ZVJoeUo0NXgzQVNQN2NPQithYW85MGZjcHhTdi9FWkZibmlBYk5nWkdoSWhwSW80SDZNSUgzTUJJR0ExVWRFd0VCL3dRSU1BWUJBZjhDQVFBd0h3WURWUjBqQkJnd0ZvQVV1N0Rlb1ZnemlKcWtpcG5ldnIzcnI5ckxKS3N3UmdZSUt3WUJCUVVIQVFFRU9qQTRNRFlHQ0NzR0FRVUZCekFCaGlwb2RIUndPaTh2YjJOemNDNWhjSEJzWlM1amIyMHZiMk56Y0RBekxXRndjR3hsY205dmRHTmhaek13TndZRFZSMGZCREF3TGpBc29DcWdLSVltYUhSMGNEb3ZMMk55YkM1aGNIQnNaUzVqYjIwdllYQndiR1Z5YjI5MFkyRm5NeTVqY213d0hRWURWUjBPQkJZRUZEOHZsQ05SMDFESm1pZzk3YkI4NWMrbGtHS1pNQTRHQTFVZER3RUIvd1FFQXdJQkJqQVFCZ29xaGtpRzkyTmtCZ0lCQkFJRkFEQUtCZ2dxaGtqT1BRUURBd05vQURCbEFqQkFYaFNxNUl5S29nTUNQdHc0OTBCYUI2NzdDYUVHSlh1ZlFCL0VxWkdkNkNTamlDdE9udU1UYlhWWG14eGN4ZmtDTVFEVFNQeGFyWlh2TnJreFUzVGtVTUkzM3l6dkZWVlJUNHd4V0pDOTk0T3NkY1o0K1JHTnNZRHlSNWdtZHIwbkRHZz0iLCJNSUlDUXpDQ0FjbWdBd0lCQWdJSUxjWDhpTkxGUzVVd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NVFF3TkRNd01UZ3hPVEEyV2hjTk16a3dORE13TVRneE9UQTJXakJuTVJzd0dRWURWUVFEREJKQmNIQnNaU0JTYjI5MElFTkJJQzBnUnpNeEpqQWtCZ05WQkFzTUhVRndjR3hsSUVObGNuUnBabWxqWVhScGIyNGdRWFYwYUc5eWFYUjVNUk13RVFZRFZRUUtEQXBCY0hCc1pTQkpibU11TVFzd0NRWURWUVFHRXdKVlV6QjJNQkFHQnlxR1NNNDlBZ0VHQlN1QkJBQWlBMklBQkpqcEx6MUFjcVR0a3lKeWdSTWMzUkNWOGNXalRuSGNGQmJaRHVXbUJTcDNaSHRmVGpqVHV4eEV0WC8xSDdZeVlsM0o2WVJiVHpCUEVWb0EvVmhZREtYMUR5eE5CMGNUZGRxWGw1ZHZNVnp0SzUxN0lEdll1VlRaWHBta09sRUtNYU5DTUVBd0hRWURWUjBPQkJZRUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdEZ1lEVlIwUEFRSC9CQVFEQWdFR01Bb0dDQ3FHU000OUJBTURBMmdBTUdVQ01RQ0Q2Y0hFRmw0YVhUUVkyZTN2OUd3T0FFWkx1Tit5UmhIRkQvM21lb3locG12T3dnUFVuUFdUeG5TNGF0K3FJeFVDTUcxbWloREsxQTNVVDgyTlF6NjBpbU9sTTI3amJkb1h0MlFmeUZNbStZaGlkRGtMRjF2TFVhZ002QmdENTZLeUtBPT0iXX0.eyJvcmlnaW5hbFRyYW5zYWN0aW9uSWQiOiIyMDAwMDAwMDAwODQyNjA3IiwiYXV0b1JlbmV3UHJvZHVjdElkIjoiY29tLmF1ZGFvcy5hdWRhcmVjb3JkZXIudmlwLm0yIiwicHJvZHVjdElkIjoiY29tLmF1ZGFvcy5hdWRhcmVjb3JkZXIudmlwLm0yIiwiYXV0b1JlbmV3U3RhdHVzIjoxLCJzaWduZWREYXRlIjoxNjQ2Mzg3MDA4MjI4LCJlbnZpcm9ubWVudCI6IlNhbmRib3gifQ.jOzVnA3eOd8xui8_pJicu9WXVHZzZK104LMTmr_C6eLvtdhzFUQ-ZDgcB2ET_clBdG9KkpxW78Ewomlj-aFfOQ",
		        "signedTransactionInfo":"eyJhbGciOiJFUzI1NiIsIng1YyI6WyJNSUlFTURDQ0E3YWdBd0lCQWdJUWFQb1BsZHZwU29FSDBsQnJqRFB2OWpBS0JnZ3Foa2pPUFFRREF6QjFNVVF3UWdZRFZRUURERHRCY0hCc1pTQlhiM0pzWkhkcFpHVWdSR1YyWld4dmNHVnlJRkpsYkdGMGFXOXVjeUJEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURUxNQWtHQTFVRUN3d0NSell4RXpBUkJnTlZCQW9NQ2tGd2NHeGxJRWx1WXk0eEN6QUpCZ05WQkFZVEFsVlRNQjRYRFRJeE1EZ3lOVEF5TlRBek5Gb1hEVEl6TURreU5EQXlOVEF6TTFvd2daSXhRREErQmdOVkJBTU1OMUJ5YjJRZ1JVTkRJRTFoWXlCQmNIQWdVM1J2Y21VZ1lXNWtJR2xVZFc1bGN5QlRkRzl5WlNCU1pXTmxhWEIwSUZOcFoyNXBibWN4TERBcUJnTlZCQXNNSTBGd2NHeGxJRmR2Y214a2QybGtaU0JFWlhabGJHOXdaWElnVW1Wc1lYUnBiMjV6TVJNd0VRWURWUVFLREFwQmNIQnNaU0JKYm1NdU1Rc3dDUVlEVlFRR0V3SlZVekJaTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEEwSUFCT29UY2FQY3BlaXBOTDllUTA2dEN1N3BVY3dkQ1hkTjh2R3FhVWpkNThaOHRMeGlVQzBkQmVBK2V1TVlnZ2gxLzVpQWsrRk14VUZtQTJhMXI0YUNaOFNqZ2dJSU1JSUNCREFNQmdOVkhSTUJBZjhFQWpBQU1COEdBMVVkSXdRWU1CYUFGRDh2bENOUjAxREptaWc5N2JCODVjK2xrR0taTUhBR0NDc0dBUVVGQndFQkJHUXdZakF0QmdnckJnRUZCUWN3QW9ZaGFIUjBjRG92TDJObGNuUnpMbUZ3Y0d4bExtTnZiUzkzZDJSeVp6WXVaR1Z5TURFR0NDc0dBUVVGQnpBQmhpVm9kSFJ3T2k4dmIyTnpjQzVoY0hCc1pTNWpiMjB2YjJOemNEQXpMWGQzWkhKbk5qQXlNSUlCSGdZRFZSMGdCSUlCRlRDQ0FSRXdnZ0VOQmdvcWhraUc5Mk5rQlFZQk1JSCtNSUhEQmdnckJnRUZCUWNDQWpDQnRneUJzMUpsYkdsaGJtTmxJRzl1SUhSb2FYTWdZMlZ5ZEdsbWFXTmhkR1VnWW5rZ1lXNTVJSEJoY25SNUlHRnpjM1Z0WlhNZ1lXTmpaWEIwWVc1alpTQnZaaUIwYUdVZ2RHaGxiaUJoY0hCc2FXTmhZbXhsSUhOMFlXNWtZWEprSUhSbGNtMXpJR0Z1WkNCamIyNWthWFJwYjI1eklHOW1JSFZ6WlN3Z1kyVnlkR2xtYVdOaGRHVWdjRzlzYVdONUlHRnVaQ0JqWlhKMGFXWnBZMkYwYVc5dUlIQnlZV04wYVdObElITjBZWFJsYldWdWRITXVNRFlHQ0NzR0FRVUZCd0lCRmlwb2RIUndPaTh2ZDNkM0xtRndjR3hsTG1OdmJTOWpaWEowYVdacFkyRjBaV0YxZEdodmNtbDBlUzh3SFFZRFZSME9CQllFRkNPQ21NQnEvLzFMNWltdlZtcVgxb0NZZXFyTU1BNEdBMVVkRHdFQi93UUVBd0lIZ0RBUUJnb3Foa2lHOTJOa0Jnc0JCQUlGQURBS0JnZ3Foa2pPUFFRREF3Tm9BREJsQWpFQWw0SkI5R0pIaXhQMm51aWJ5VTFrM3dyaTVwc0dJeFBNRTA1c0ZLcTdoUXV6dmJleUJ1ODJGb3p6eG1ienBvZ29BakJMU0ZsMGRaV0lZbDJlalBWK0RpNWZCbktQdThteW1CUXRvRS9IMmJFUzBxQXM4Yk51ZVUzQ0JqamgxbHduRHNJPSIsIk1JSURGakNDQXB5Z0F3SUJBZ0lVSXNHaFJ3cDBjMm52VTRZU3ljYWZQVGp6Yk5jd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NakV3TXpFM01qQXpOekV3V2hjTk16WXdNekU1TURBd01EQXdXakIxTVVRd1FnWURWUVFERER0QmNIQnNaU0JYYjNKc1pIZHBaR1VnUkdWMlpXeHZjR1Z5SUZKbGJHRjBhVzl1Y3lCRFpYSjBhV1pwWTJGMGFXOXVJRUYxZEdodmNtbDBlVEVMTUFrR0ExVUVDd3dDUnpZeEV6QVJCZ05WQkFvTUNrRndjR3hsSUVsdVl5NHhDekFKQmdOVkJBWVRBbFZUTUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUVic1FLQzk0UHJsV21aWG5YZ3R4emRWSkw4VDBTR1luZ0RSR3BuZ24zTjZQVDhKTUViN0ZEaTRiQm1QaENuWjMvc3E2UEYvY0djS1hXc0w1dk90ZVJoeUo0NXgzQVNQN2NPQithYW85MGZjcHhTdi9FWkZibmlBYk5nWkdoSWhwSW80SDZNSUgzTUJJR0ExVWRFd0VCL3dRSU1BWUJBZjhDQVFBd0h3WURWUjBqQkJnd0ZvQVV1N0Rlb1ZnemlKcWtpcG5ldnIzcnI5ckxKS3N3UmdZSUt3WUJCUVVIQVFFRU9qQTRNRFlHQ0NzR0FRVUZCekFCaGlwb2RIUndPaTh2YjJOemNDNWhjSEJzWlM1amIyMHZiMk56Y0RBekxXRndjR3hsY205dmRHTmhaek13TndZRFZSMGZCREF3TGpBc29DcWdLSVltYUhSMGNEb3ZMMk55YkM1aGNIQnNaUzVqYjIwdllYQndiR1Z5YjI5MFkyRm5NeTVqY213d0hRWURWUjBPQkJZRUZEOHZsQ05SMDFESm1pZzk3YkI4NWMrbGtHS1pNQTRHQTFVZER3RUIvd1FFQXdJQkJqQVFCZ29xaGtpRzkyTmtCZ0lCQkFJRkFEQUtCZ2dxaGtqT1BRUURBd05vQURCbEFqQkFYaFNxNUl5S29nTUNQdHc0OTBCYUI2NzdDYUVHSlh1ZlFCL0VxWkdkNkNTamlDdE9udU1UYlhWWG14eGN4ZmtDTVFEVFNQeGFyWlh2TnJreFUzVGtVTUkzM3l6dkZWVlJUNHd4V0pDOTk0T3NkY1o0K1JHTnNZRHlSNWdtZHIwbkRHZz0iLCJNSUlDUXpDQ0FjbWdBd0lCQWdJSUxjWDhpTkxGUzVVd0NnWUlLb1pJemowRUF3TXdaekViTUJrR0ExVUVBd3dTUVhCd2JHVWdVbTl2ZENCRFFTQXRJRWN6TVNZd0pBWURWUVFMREIxQmNIQnNaU0JEWlhKMGFXWnBZMkYwYVc5dUlFRjFkR2h2Y21sMGVURVRNQkVHQTFVRUNnd0tRWEJ3YkdVZ1NXNWpMakVMTUFrR0ExVUVCaE1DVlZNd0hoY05NVFF3TkRNd01UZ3hPVEEyV2hjTk16a3dORE13TVRneE9UQTJXakJuTVJzd0dRWURWUVFEREJKQmNIQnNaU0JTYjI5MElFTkJJQzBnUnpNeEpqQWtCZ05WQkFzTUhVRndjR3hsSUVObGNuUnBabWxqWVhScGIyNGdRWFYwYUc5eWFYUjVNUk13RVFZRFZRUUtEQXBCY0hCc1pTQkpibU11TVFzd0NRWURWUVFHRXdKVlV6QjJNQkFHQnlxR1NNNDlBZ0VHQlN1QkJBQWlBMklBQkpqcEx6MUFjcVR0a3lKeWdSTWMzUkNWOGNXalRuSGNGQmJaRHVXbUJTcDNaSHRmVGpqVHV4eEV0WC8xSDdZeVlsM0o2WVJiVHpCUEVWb0EvVmhZREtYMUR5eE5CMGNUZGRxWGw1ZHZNVnp0SzUxN0lEdll1VlRaWHBta09sRUtNYU5DTUVBd0hRWURWUjBPQkJZRUZMdXczcUZZTTRpYXBJcVozcjY5NjYvYXl5U3JNQThHQTFVZEV3RUIvd1FGTUFNQkFmOHdEZ1lEVlIwUEFRSC9CQVFEQWdFR01Bb0dDQ3FHU000OUJBTURBMmdBTUdVQ01RQ0Q2Y0hFRmw0YVhUUVkyZTN2OUd3T0FFWkx1Tit5UmhIRkQvM21lb3locG12T3dnUFVuUFdUeG5TNGF0K3FJeFVDTUcxbWloREsxQTNVVDgyTlF6NjBpbU9sTTI3amJkb1h0MlFmeUZNbStZaGlkRGtMRjF2TFVhZ002QmdENTZLeUtBPT0iXX0.eyJ0cmFuc2FjdGlvbklkIjoiMjAwMDAwMDAwNDA0NzExOSIsIm9yaWdpbmFsVHJhbnNhY3Rpb25JZCI6IjIwMDAwMDAwMDA4NDI2MDciLCJ3ZWJPcmRlckxpbmVJdGVtSWQiOiIyMDAwMDAwMDAwMzAyODMyIiwiYnVuZGxlSWQiOiJjb20uYXVkYW9zLmF1ZGFyZWNvcmRlciIsInByb2R1Y3RJZCI6ImNvbS5hdWRhb3MuYXVkYXJlY29yZGVyLnZpcC5tMiIsInN1YnNjcmlwdGlvbkdyb3VwSWRlbnRpZmllciI6IjIwOTI5NTM2IiwicHVyY2hhc2VEYXRlIjoxNjQ2Mzg3MDE2MDAwLCJvcmlnaW5hbFB1cmNoYXNlRGF0ZSI6MTY0NjA0NjAzNzAwMCwiZXhwaXJlc0RhdGUiOjE2NDYzODcxOTYwMDAsInF1YW50aXR5IjoxLCJ0eXBlIjoiQXV0by1SZW5ld2FibGUgU3Vic2NyaXB0aW9uIiwiaW5BcHBPd25lcnNoaXBUeXBlIjoiUFVSQ0hBU0VEIiwic2lnbmVkRGF0ZSI6MTY0NjM4NzAwODI1NCwiZW52aXJvbm1lbnQiOiJTYW5kYm94In0.mvVlACDCcS-l1oAYRI9JVsy3dOQ0snjdA7Qlag15DrQ8lmAosrkjrBRPzOnTCaNUkUgusOMdBDDPfhfoSk0oMw"
		    }
		}
	*/

	// decode renewalInfo
	renewalInfo, err := payload.DecodeRenewalInfo()
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("data.renewalInfo: %+v", renewalInfo)
	bs, _ := json.Marshal(renewalInfo)
	xlog.Color(xlog.GreenBright).Info(string(bs))
	/*
		{
		    "autoRenewProductId":"com.audaos.audarecorder.vip.m2",
		    "autoRenewStatus":1,
		    "expirationIntent":0,
		    "gracePeriodExpiresDate":0,
		    "isInBillingRetryPeriod":false,
		    "offerIdentifier":"",
		    "offerType":0,
		    "originalTransactionId":"2000000000842607",
		    "priceIncreaseStatus":0,
		    "productId":"com.audaos.audarecorder.vip.m2",
		    "signedDate":1646387008228
		}
	*/
	// decode transactionInfo
	transactionInfo, err := payload.DecodeTransactionInfo()
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("data.transactionInfo: %+v", transactionInfo)
	bs2, _ := json.Marshal(transactionInfo)
	xlog.Color(xlog.YellowBright).Info(string(bs2))
	/*
		{
		    "appAccountToken":"",
		    "bundleId":"com.audaos.audarecorder",
		    "expiresDate":1646387196000,
		    "inAppOwnershipType":"PURCHASED",
		    "isUpgraded":false,
		    "offerIdentifier":"",
		    "offerType":0,
		    "originalPurchaseDate":1646046037000,
		    "originalTransactionId":"2000000000842607",
		    "productId":"com.audaos.audarecorder.vip.m2",
		    "purchaseDate":1646387016000,
		    "quantity":1,
		    "revocationDate":0,
		    "revocationReason":"",
		    "signedDate":1646387008254,
		    "subscriptionGroupIdentifier":"20929536",
		    "transactionId":"2000000004047119",
		    "type":"Auto-Renewable Subscription",
		    "webOrderLineItemId":"2000000000302832"
		}
	*/
}
