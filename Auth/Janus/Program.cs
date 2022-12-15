using System.Data.Entity;
using Janus;
using Janus.Identity;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using DbContext = Janus.DbContext;

var config = new ConfigurationBuilder().AddJsonFile("appsettings.json", optional: false).Build();

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();
builder.Services.AddDbContext<DbContext>(options => options.UseNpgsql(Environment.GetEnvironmentVariable("SqlConnection") ?? config.GetConnectionString("SqlConnection")));
builder.Services.AddIdentity<User, IdentityRole>(options =>
{
    options.SignIn.RequireConfirmedAccount = false;
    options.Lockout.AllowedForNewUsers = true;
    options.Lockout.DefaultLockoutTimeSpan = TimeSpan.FromMinutes(5);
    options.Lockout.MaxFailedAccessAttempts = 5;
}).AddEntityFrameworkStores<DbContext>().AddRoles<IdentityRole>();

builder.Services.Configure<IdentityOptions>(options =>
{
    options.User.RequireUniqueEmail = true;
});

builder.Services.AddCors(options =>
{
    options.AddPolicy("AllowAll", builder =>
    {
        // builder.WithOrigins("http://localhost:3000")
        //     .AllowAnyHeader()
        //     .AllowAnyMethod()
        //     .AllowCredentials();
        // builder.WithOrigins("http://localhost")
        //     .AllowAnyHeader()
        //     .AllowAnyMethod()
        //     .AllowCredentials();
        builder.WithOrigins("http://patrickbuck.net")
            .AllowAnyHeader()
            .AllowAnyMethod()
            .AllowCredentials();
        builder.WithOrigins("https://patrickbuck.net")
            .AllowAnyHeader()
            .AllowAnyMethod()
            .AllowCredentials();
    });
});

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseCors("AllowAll");

app.UseHttpsRedirection();

app.MapControllers();

app.Run();
