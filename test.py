
from requests import get

# 1234 just listen on localip
print(get('https://meuip.com/api/meuip.php', proxies={'https':'socks5://13.211.239.12:1234'}).text)
