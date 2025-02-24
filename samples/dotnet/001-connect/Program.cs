IConfiguration configuration = new ConfigurationBuilder()
    .AddEnvironmentVariables()
    .AddUserSecrets<Program>()
    .Build();

string credential = configuration["SQL_CREDENTIAL"] ??
    throw new InvalidOperationException("Please set the .NET user secret (or environment variable) \"SQL_CREDENTIAL\" using \"dotnet user-secrets set\"");

Console.MarkupLine("[blue bold]:gear:\tConnecting to SQL Server...[/]");

using SqlConnection connection = new(credential);

const string query = "SELECT @@VERSION";

Console.MarkupLine($"[blue]Query:\t[underline]{query}[/][/]");

using SqlCommand command = new(query, connection);

await connection.OpenAsync();

string? result = await command.ExecuteScalarAsync() as string;

if (result is not null)
    Console.MarkupLine($"[green]Result:\t{result}[/]");