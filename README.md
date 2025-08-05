# PortFWD 🔀

**PortFWD** es una aplicación sencilla escrita en **Go** que permite realizar **port forwarding** de manera local o remota entre puertos TCP. Pensada para entornos donde no es posible usar SSH, VPN o herramientas como `chisel`, esta herramienta ayuda a exponer puertos internos de servicios (como HTTP, FTP, bases de datos) hacia el exterior sin necesidad de privilegios elevados.

---

## 🛠️ ¿Para qué sirve PortFWD?

Las aplicaciones de port forwarding como PortFWD son esenciales en escenarios como:

- **Pentesting / CTFs** 🕵️‍♂️: Redirigir puertos locales de máquinas objetivo hacia el atacante para escaneo.
- **Infraestructura restringida** 🔐: Acceder a servicios que solo corren en localhost sin acceso directo.
- **Seguridad ofensiva / evasión** 🚪: Evitar controles o restricciones que impiden exponer servicios.
- **Desarrollo de redes** 📡: Reenviar puertos entre entornos aislados sin necesidad de configuración pesada.

---

## ⚙️ Instalación

Clona este repositorio:

```bash
git clone https://github.com/aratan/PortFWD.git
cd PortFWD
go build -o portfwd portfwd.go
```
uso:
```bash
./portfwd -lport 8000 -rhost 127.0.0.1 -rport 80
```
