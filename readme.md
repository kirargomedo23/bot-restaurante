# Bot Restaurante

Este proyecto es un bot de restaurante desarrollado en Go, que utiliza Firestore como base de datos y Gemini AI para responder preguntas sobre el menú. El bot permite consultar información de platos, calcular precios y responder preguntas frecuentes.

## Características
 
- Integración con Firestore para almacenar y consultar el menú.
- Uso de Gemini AI para respuestas inteligentes.
- Variables de entorno gestionadas con `.env`. 

## Configuración

1. **Clona el repositorio:**
   ```sh
   git clone https://github.com/kirargomedo23/bot-restaurante.git
   cd bot-restaurante
   ```

2. **Configura las variables de entorno:**
   - Crea un archivo `.env` en la raíz del proyecto y añade tus credenciales:
   ```env
   API_KEY_GEMINI=tu_api_key_gemini
   GOOGLE_APPLICATION_CREDENTIALS=path/a/tu/credencial.json
   PROJECT_ID=tu_project_id
   ```

   - Crea un archivo serviceAccountKey.json en la raíz del proyecto
    


3. **Instala las dependencias:**
   Asegúrate de tener Go instalado y ejecuta:
   ```sh
   go mod tidy
   ```

4. **Ejecutar el bot:**
   ```sh
   make run 
   ``


